package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type configuration struct {
	Port         string `json:"port"`
	SSL          bool   `json:"ssl"`
	Cert         string `json:"cert,omitempty"`
	Key          string `json:"key,omitempty"`
	Database     string `json:"database"`
	AuthProvider string `json:"authprovider"`
}

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	debugMode     bool
}

func main() {
	//dsn := flag.String("dsn", "user:pass@/dbName?parseTime=true", "MySQL data source name")
	dbg := flag.Bool("debug", false, "enable debug mode")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := &configuration{}

	cfgReader, err := os.ReadFile("apotheca.json")
	if err != nil {
		logger.Error("failed to read configuration json", slog.String("error", err.Error()))
	}

	err = json.Unmarshal(cfgReader, cfg)
	if err != nil {
		logger.Error("failed to unmarshal configuration json", slog.String("error", err.Error()))
	}

	logger.Info("Configuration", slog.Any("config", cfg))

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
		Addr:         cfg.Port,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server on", slog.String("addr", srv.Addr))

	if cfg.SSL {
		err = srv.ListenAndServeTLS(cfg.Cert, cfg.Key)
		logger.Error("failed to start SSL server", slog.String("error", err.Error()))
		os.Exit(1)
	}
	err = srv.ListenAndServe()
	logger.Error("failed to start server", slog.String("error", err.Error()))
	os.Exit(1)
}
