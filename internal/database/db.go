package database

import (
	"log"

	"github.com/benjaminrae/authentication/internal/database/models"
	env "github.com/benjaminrae/authentication/internal/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := env.Config.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&models.Tenant{})
}
