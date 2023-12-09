package server

import (
	"net/http"

	"github.com/benjaminrae/authentication/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	router := echo.New()

	router.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	router.POST("/tenants", handlers.CreateTenantHandler())

	router.GET("/tenants", handlers.GetTenantsHandler())

	return router
}
