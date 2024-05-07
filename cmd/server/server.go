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
}

func RunServer() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := &Configuration{}

	cfgReader, err := os.ReadFile("apotheca.json")
	if err != nil {
		logger.Error("failed to read configuration json", slog.String("error", err.Error()))
	}

	err = json.Unmarshal(cfgReader, cfg)
	if err != nil {
		logger.Error("failed to unmarshal configuration json", slog.String("error", err.Error()))
	}

	logger.Info("Configuration", slog.Any("config", cfg))

	mcit, err := time.ParseDuration(cfg.Database.DBMaxIdleTime)
	if err != nil {
		logger.Error("failed to parse database max idle time", slog.String("error", err.Error()))
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
	logger.Info("Data Source", slog.String("dsn", dsn))

	db, err := openDB(dsn)
	if err != nil {
		logger.Error("Database", slog.String("error", err.Error()))
		db.Close()
		os.Exit(1)
	}
	defer db.Close()

	tc, err := newTemplateCache()
	if err != nil {
		logger.Error("failed to initialize template cache", slog.String("error", err.Error()))
	}

	app := &application{
		logger:        logger,
		templateCache: tc,
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
		tlsConfig := &tls.Config{
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		}
		srv.TLSConfig = tlsConfig
		err = srv.ListenAndServeTLS(cfg.Cert, cfg.Key)
		logger.Error("failed to start SSL server", slog.String("error", err.Error()))
		os.Exit(1)
	}
	err = srv.ListenAndServe()
	logger.Error("failed to start server", slog.String("error", err.Error()))
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
