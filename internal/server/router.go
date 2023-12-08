package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	router := echo.New()

	router.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	return router
}
