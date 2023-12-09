package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/benjaminrae/authentication/internal/database"
	"github.com/benjaminrae/authentication/internal/database/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateTenantRequest struct {
	Name string `json:"name"`
}

type CreateTenantResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
}

func CreateTenantHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Received CreateTenant request")
		var request CreateTenantRequest

		err := json.NewDecoder(c.Request().Body).Decode(&request)

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, fmt.Sprintf("Bad request: %v", err))
		}

		tenant, err := models.CreateTenant(request.Name)

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Couldn't create Tenant: %v", err))
		}

		if err := database.DB.Create(&tenant).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Couldn't create Tenant: %v", err))
		}

		response := CreateTenantResponse{
			ID:        tenant.ID,
			CreatedAt: tenant.CreatedAt,
			UpdatedAt: tenant.UpdatedAt,
			DeletedAt: tenant.UpdatedAt,
			Name:      tenant.Name,
			Key:       tenant.Key,
		}

		return c.JSON(http.StatusCreated, &response)
	}
}
