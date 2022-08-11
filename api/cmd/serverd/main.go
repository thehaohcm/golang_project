package main

import (
	"golang_project/api/cmd/serverd/database"
	"golang_project/api/cmd/serverd/routes"
)

func main() {
	server := routes.SetupRouter()
	server.Run(":8080")

	defer database.CloseDB()
}
