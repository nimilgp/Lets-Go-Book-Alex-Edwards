package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", "3333", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to static assets")
	flag.Parse()

	var app application
	app.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	app.logger.Info("starting server at port:", "addr", cfg.addr)

	if err := http.ListenAndServe(":"+cfg.addr, app.route(cfg)); err != nil {
		app.logger.Error(err.Error())
		os.Exit(-1)
	}
}
