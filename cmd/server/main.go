package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/misham/appointment-scheduler/pkg/calendar"
	"github.com/misham/appointment-scheduler/pkg/version"
	"go.uber.org/zap"
)

func main() {
	addr := flag.String("host", "127.0.0.1", "HTTP network address")
	port := flag.String("port", "8080", "HTTP network port")
	showVersion := flag.Bool("version", false, "Show version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println("Version:", version.Version)
		os.Exit(0)
	}

	url := fmt.Sprintf("%s:%s", *addr, *port)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	errorLog, err := zap.NewStdLogAt(logger, zap.ErrorLevel)
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	sugar := logger.Sugar()

	app := &Server{
		logger:    sugar,
		errorLog:  errorLog,
		calendars: calendar.NewCalendarModel(),
	}

	srv := &http.Server{
		Addr:         url,
		Handler:      app.routes(),
		ErrorLog:     app.errorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// TODO handle graceful shutdown
	sugar.Info("starting server", "addr", srv.Addr)
	err = srv.ListenAndServe()
	sugar.Error(err.Error())
	os.Exit(1)
}
