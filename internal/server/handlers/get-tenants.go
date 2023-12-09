package handlers

import (
	"fmt"
	"net/http"

	"github.com/benjaminrae/authentication/internal/database"
	"github.com/benjaminrae/authentication/internal/database/models"
	"github.com/labstack/echo/v4"
)

func GetTenantsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tenants []models.Tenant

		if err := database.DB.Find(&tenants).Error; err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error fetching tenants: %v", err))
		}

		return c.JSON(http.StatusOK, &tenants)
	}
}
