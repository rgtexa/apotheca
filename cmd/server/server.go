package server

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Configuration struct {
	Port     string `json:"port"`
	SSL      bool   `json:"ssl"`
	Cert     string `json:"cert,omitempty"`
	Key      string `json:"key,omitempty"`
	Database struct {
		DBProvider     string `json:"dbprovider"`
		DBHost         string `json:"dbhost"`
		DBPort         string `json:"dbport"`
		DBUser         string `json:"dbuser"`
		DBPass         string `json:"dbpass"`
		DBName         string `json:"dbname"`
		DBMaxOpenConns int    `json:"dbmaxopenconns"`
		DBMaxIdleTime  string `json:"dbmaxidletime"`
	} `json:"database"`
	AuthProvider string `json:"authprovider"`
}

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	debug         bool
}

func RunServer(runFlags []string) {

	logLevel := new(slog.LevelVar)
	logOptions := slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.NewJSONHandler(os.Stdout, &logOptions)
	slog.SetDefault(slog.New(logger))

	tc, err := newTemplateCache()
	if err != nil {
		slog.Error("failed to initialize template cache", slog.String("error", err.Error()))
	}

	app := &application{
		logger:        &slog.Logger{},
		templateCache: tc,
	}

	if len(runFlags) > 0 {
		app.debug = true
		logLevel.Set(slog.LevelDebug)
		slog.Debug("Run Flags", slog.String("Flags", runFlags[0]))
	}

	cfg := &Configuration{}

	cfgReader, err := os.ReadFile("apotheca.json")
	if err != nil {
		slog.Error("failed to read configuration json", slog.String("error", err.Error()))
	}

	err = json.Unmarshal(cfgReader, cfg)
	if err != nil {
		slog.Error("failed to unmarshal configuration json", slog.String("error", err.Error()))
	}

	slog.Info("Configuration", slog.Any("config", cfg))

	mcit, err := time.ParseDuration(cfg.Database.DBMaxIdleTime)
	if err != nil {
		slog.Error("failed to parse database max idle time", slog.String("error", err.Error()))
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?pool_max_conns=%d&pool_max_conn_idle_time=%v",
		cfg.Database.DBProvider,
		cfg.Database.DBUser,
		cfg.Database.DBPass,
		cfg.Database.DBHost,
		cfg.Database.DBPort,
		cfg.Database.DBName,
		cfg.Database.DBMaxOpenConns,
		mcit,
	)
	slog.Debug("Data Source", slog.String("dsn", dsn))

	db, err := openDB(dsn)
	if err != nil {
		slog.Error("Database", slog.String("error", err.Error()))
		db.Close()
		os.Exit(1)
	}
	defer db.Close()

	srv := &http.Server{
		Addr:         cfg.Port,
		ErrorLog:     slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("starting server on", slog.String("addr", srv.Addr))

	if cfg.SSL {
		tlsConfig := &tls.Config{
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		}
		srv.TLSConfig = tlsConfig
		err = srv.ListenAndServeTLS(cfg.Cert, cfg.Key)
		slog.Error("failed to start SSL server", slog.String("error", err.Error()))
		os.Exit(1)
	}
	err = srv.ListenAndServe()
	slog.Error("failed to start server", slog.String("error", err.Error()))
	os.Exit(1)
}

func openDB(dsn string) (*pgxpool.Pool, error) {

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = conn.Ping(ctx)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}
