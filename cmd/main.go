package main

import (
	"github.com/benjaminrae/authentication/internal/database"
	env "github.com/benjaminrae/authentication/internal/environment"
	"github.com/benjaminrae/authentication/internal/server"
)

func main() {
	env.Load()

	server := server.New()

	database.Init()

	server.Logger.Fatal(server.Start(env.Config.Port))

}
