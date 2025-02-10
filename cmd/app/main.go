package main

import (
	config "ex_proj_go/configs"
	"ex_proj_go/internal/db"

	"log"
)

// @title ExProjectGo API
// @version 1.0
// @description This is a sample API with Gin, GORM, Swagger, and Worker Pool.

// @host localhost:8080
// @BasePath /api/v1

func main() {

	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}

	// Initialize database
	postgresDB, err := db.InitPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

}
