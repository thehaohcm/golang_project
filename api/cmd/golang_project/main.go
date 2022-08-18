package main

import (
	"golang_project/api/internal/api/router"
	"golang_project/api/internal/config"
)

func main() {
	server := router.SetupRouter()
	server.Run(":8080")

	defer config.CloseDB()
}
