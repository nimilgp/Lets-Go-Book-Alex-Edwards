package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nimilgp/paste-bin/internal/models"
)

type config struct {
	addr      string
	staticDir string
	dsn       string
}

type application struct {
	logger   *slog.Logger
	snippets *models.SnipetModel
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", "3333", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to static assets")
	flag.StringVar(&cfg.dsn, "dsn", "web:super-secret-passwd@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	var app application
	app.logger = logger
	app.snippets = &models.SnipetModel{DB: db}

	app.logger.Info("starting server at port:", "addr", cfg.addr)
	if err := http.ListenAndServe(":"+cfg.addr, app.route(cfg)); err != nil {
		app.logger.Error(err.Error())
		os.Exit(-1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
