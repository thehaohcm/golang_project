package main

import (
	"os"

	"golang_project/api/internal/api/router"
	"golang_project/api/internal/config"
)

func main() {
	server := router.SetupRouter()
	port := ":" + os.Getenv("APP_PORT")
	server.Run(port)

	defer config.CloseDB()
}
