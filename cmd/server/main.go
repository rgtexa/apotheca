package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	debugMode     bool
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	//dsn := flag.String("dsn", "user:pass@/dbName?parseTime=true", "MySQL data source name")
	dbg := flag.Bool("debug", false, "enable debug mode")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	tc, err := newTemplateCache()
	if err != nil {
		logger.Error("failed to initialize template cache", slog.String("error", err.Error()))
	}

	app := &application{
		logger:        logger,
		templateCache: tc,
		debugMode:     *dbg,
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server on", slog.String("addr", srv.Addr))

	err = srv.ListenAndServe()
	logger.Error("failed to start server", slog.String("error", err.Error()))
	os.Exit(1)
}
