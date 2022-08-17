package main

import (
	"golang_project/api/cmd/golang_project/database"
	"golang_project/api/internal/api/router"
)

func main() {
	server := router.SetupRouter()
	server.Run(":8080")

	defer database.CloseDB()
}
