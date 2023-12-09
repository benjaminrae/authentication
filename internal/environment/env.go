package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Port string
	Dsn  string
}

var Config Environment

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	rawPort := os.Getenv("PORT")
	port := ":" + rawPort

	dsn := os.Getenv("POSTGRES_DSN")

	Config = Environment{
		Port: port,
		Dsn:  dsn,
	}
}
