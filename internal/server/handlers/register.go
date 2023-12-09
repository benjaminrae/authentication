package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benjaminrae/authentication/internal/database"
	"github.com/benjaminrae/authentication/internal/database/models"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	FirstName string
	LastName  string
	Password  string
}

func RegisterHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var newUser models.User
		err := json.NewDecoder(c.Request().Body).Decode(&newUser)

		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		}

		if err := database.DB.Create(&newUser).Error; err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
		}

		return c.JSON(http.StatusCreated, &newUser)
	}
}
