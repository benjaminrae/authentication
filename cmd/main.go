package main

import (
	"github.com/benjaminrae/authentication/internal/database"
	"github.com/benjaminrae/authentication/internal/server"
)

func main() {
	server := server.New()

	database.Init()

	server.Logger.Fatal(server.Start(":4000"))

}
