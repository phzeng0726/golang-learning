package main

import (
	"fmt"
	"learning/pkg/config"
	"learning/pkg/handlers"
	"learning/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// move to outside because the NoSurf function used this variable
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this value to true when the env is production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	handlers.NewHandlers(&app)
	render.NewTemplates(&app)

	fmt.Printf("Server port number: %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
