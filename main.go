package main

import (
	"bibliosphere_gin/adapters/database"
	"bibliosphere_gin/cmd"
	"bibliosphere_gin/config"
	"log"
)

func main() {
	config.LoadConfig()
	cmd.Execute()

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := setupRouter(db)
	port := config.AppConfig.DefaultPort
	router.Run(port)
}
