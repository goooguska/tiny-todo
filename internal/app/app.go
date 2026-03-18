package app

import (
	"errors"
	"net/http"
	"time"
	"tiny-todo/internal/http/router"
)

type App struct {
	diContainer *diContainer
	httpServer  *http.Server
}

func New() *App {
	a := &App{}

	a.Bootstrap()

	return a
}

func (a *App) Run() {
	defer a.diContainer.Db().DB.Close()
	a.diContainer.Logger().Info("starting server", "addr", a.diContainer.Config().Server.Port)
	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.diContainer.Logger().Error("start server failed", "error", err)
	}
}

func (a *App) Bootstrap() *App {
	a.initDI()
	a.initHttpServer()

	return a
}

func (a *App) initDI() {
	a.diContainer = NewDiContainer()
}

func (a *App) initHttpServer() {
	r := router.NewRouter(a.diContainer.ApiHandlers())

	a.httpServer = &http.Server{
		Addr:        ":" + a.diContainer.Config().Server.Port,
		ReadTimeout: time.Duration(a.diContainer.Config().Server.Timeout) * time.Minute,
		Handler:     r,
	}
}
