package main

import (
	"golang_project/api/cmd/serverd/routes"
)

func main() {
	server := routes.SetupRouter()
	server.Run(":8080")
}
