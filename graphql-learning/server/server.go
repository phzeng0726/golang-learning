package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"graphql-learning/autogen"
	"graphql-learning/resolvers"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	if enablePlayground {
		mux.Handle("/", func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/" {
					http.NotFound(w, r)
					return
				}
				next.ServeHTTP(w, r)
			})
		}(playground.Handler("programName", "/graphql")))
	}

	var dial string
	if containerMode {
		dial = fmt.Sprintf("unix(/cloudsql/%s)", instanceConnectionName)
	} else {
		dial = localDBDial
	}
	log.Printf("DB dial = %s\n", dial)

	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=UTC&timeout=%s",
		dbUser, dbPass, dbName, dbDialTimeout)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.New(log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				SlowThreshold:             dbSlowThreshold,
				Colorful:                  !containerMode,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				LogLevel:                  logLevel,
			}),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	db, err := conn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	db.SetConnMaxIdleTime(dbConnMaxIdleTime)
	db.SetConnMaxLifetime(dbConnMaxLifetime)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetMaxOpenConns(dbMaxOpenConns)

	hs := handler.New(autogen.NewExecutableSchema(resolvers.NewConfig()))
	hs.AddTransport(transport.POST{})
	hs.AddTransport(transport.MultipartForm{MaxMemory: maxMemory, MaxUploadSize: maxUploadSize})
	if enablePlayground {
		hs.Use(extension.Introspection{})
	}
	hs.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		if systemMaintenance {
			return func(ctx context.Context) *graphql.Response {
				return graphql.ErrorResponse(ctx, "system is under maintenance")
			}
		}
		// DB goroutine-safe
		return next(context.WithValue(ctx, resolvers.ContextKeyDB, conn.WithContext(ctx)))
	})

	mux.Handle("/graphql", func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// x-api-key
			if !disableAuth && r.Header.Get(keyXApiKey) != apiKey {
				http.Error(w, "403 forbidden", http.StatusForbidden)
				return
			}
			// x-api-auth
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), resolvers.ContextKeyAuth,
				r.Header.Get(keyXApiAuth))))
		})
	}(hs))

	// https://cloud.google.com/run/docs/samples/cloudrun-sigterm-handler
	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: mux}

	// SIGINT handles Ctrl+C locally
	// SIGTERM handles Cloud Run termination signal
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start HTTP server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Receive output from signalChan
	signal := <-signalChan
	log.Printf("%s signal caught\n", signal)

	// Timeout if waiting for connections to return idle
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Add extra handling here to clean up resources, such as flushing logs and
	// closing any database or Redis connections
	if err := db.Close(); err != nil {
		log.Printf("database closing failed: %+v\n", err)
	}

	// Gracefully shutdown the server by waiting on existing requests (except websockets)
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %+v\n", err)
	}
	log.Println("server exited")
}
