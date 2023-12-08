package main

import "github.com/benjaminrae/authentication/internal/server"

func main() {
	server := server.New()

	server.Logger.Fatal(server.Start(":4000"))

}
