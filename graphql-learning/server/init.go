package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gorm.io/gorm/logger"
)

const (
	apiKey                    = "mJI42RoGMWDZPKuZIQOUQ4MaMErVXrzLi0pcR56i"
	dbConnMaxIdleTime         = 3 * time.Minute
	dbConnMaxLifetime         = 6 * time.Minute
	dbDialTimeout             = "20s"
	dbMaxIdleConns            = 20
	dbMaxOpenConns            = 100
	dbSlowThreshold           = 200 * time.Millisecond
	keyDBName                 = "DB_NAME"
	keyDBPass                 = "DB_PASS"
	keyDBUser                 = "DB_USER"
	keyDisableAuth            = "DISABLE_AUTH"
	keyEnablePlayground       = "ENABLE_PLAYGROUND"
	keyInstanceConnectionName = "INSTANCE_CONNECTION_NAME"
	keyLogReduce              = "LOG_REDUCE"
	keyPort                   = "PORT"
	keyStorageBucket          = "STORAGE_BUCKET"
	keySystemMaintenance      = "SYSTEM_MAINTENANCE"
	keyXApiKey                = "x-api-key"
	keyXApiAuth               = "x-api-auth"
	localDBDial               = "tcp(35.222.182.191:3306)"
	maxMemory                 = 1 * (1 << 20)
	maxUploadSize             = 1 * (1 << 20)
	programName               = "GDC Backend GraphQL"
	version                   = "0.5.0"
)

var (
	containerMode          bool
	dbName                 string
	dbPass                 string
	dbUser                 string
	disableAuth            bool
	enablePlayground       bool
	instanceConnectionName string
	logLevel               logger.LogLevel
	port                   string
	signalChan             chan os.Signal
	storageBucket          string
	systemMaintenance      bool
)

func init() {
	flag.BoolVar(&containerMode, "c", false, "container mode.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]...\n",
			filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "\n%s v%s\n", programName, version)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}
	// dbName = os.Getenv(keyDBName)
	// dbPass = os.Getenv(keyDBPass)
	// dbUser = os.Getenv(keyDBUser)
	dbName = "graphql"
	dbPass = "0000"
	dbUser = "root"
	disableAuth, _ = strconv.ParseBool(os.Getenv(keyDisableAuth))
	enablePlayground, _ = strconv.ParseBool(os.Getenv(keyEnablePlayground))
	instanceConnectionName = os.Getenv(keyInstanceConnectionName)
	if logReduce, err := strconv.ParseBool(os.Getenv(keyLogReduce)); err == nil && logReduce {
		logLevel = logger.Warn
	} else {
		logLevel = logger.Info
	}
	port = os.Getenv(keyPort)
	if port == "" {
		port = "8080"
	}
	log.Printf("PORT = %s\n", port)
	// Create channel to listen for signals
	signalChan = make(chan os.Signal, 1)
	storageBucket = os.Getenv(keyStorageBucket)
	systemMaintenance, _ = strconv.ParseBool(os.Getenv(keySystemMaintenance))
}
