package app

import (
	"github.com/labstack/echo/v4"
	"log"
	"middleware-test/internal/app/endpoint"
	"middleware-test/internal/app/middleware"
	"middleware-test/internal/app/service"
)

type App struct {
	ep     *endpoint.Endpoint
	svc    *service.Service
	server *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.svc = service.New()
	a.ep = endpoint.New(a.svc)
	a.server = echo.New()

	a.server.Use(middleware.RoleCheck)

	a.server.GET("/", a.ep.MainPage)

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
