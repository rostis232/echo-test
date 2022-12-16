package app

import (
	"github.com/labstack/echo/v4"
	"log"
	"middleware-test/internal/app/handler"
	"middleware-test/internal/app/middleware"
	"middleware-test/internal/app/service"
)

type App struct {
	handler *handler.Handler
	service *service.Service
	server  *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.service = service.New()
	a.handler = handler.New(a.service)
	a.server = echo.New()

	a.server.Use(middleware.RoleCheck)

	a.server.GET("/", a.handler.MainPage)

	return a, nil
}

func (a *App) Run() error {
	err := a.server.Start(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server is running")
	return nil
}
