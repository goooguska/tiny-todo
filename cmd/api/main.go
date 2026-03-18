package main

import (
	"errors"
	"log"
	"net/http"
	"tiny-todo/internal/app"
	"tiny-todo/internal/config"
)

func main() {
	cfg, err := config.New("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	a := app.Bootstrap(cfg)
	defer a.DB.Close()

	server := a.HttpServer()

	a.Logger.Info("starting server", "addr", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.Logger.Error("start server failed", "error", err)
	}
}
