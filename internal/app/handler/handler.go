package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Service interface {
	Days() int64
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{
		svc,
	}
}

func (e *Handler) MainPage(ctx echo.Context) error {
	days := e.svc.Days()
	answer := fmt.Sprintf("Days: %d", days)
	err := ctx.String(http.StatusOK, answer)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
