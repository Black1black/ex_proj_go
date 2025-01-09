package db

import (
	"ex_proj_go/internal/models"
	"fmt"
	"log"

	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	configFile, err := os.Open("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}
	defer configFile.Close()

	var config map[string]map[string]string
	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config["database"]["host"], config["database"]["user"], config["database"]["password"],
		config["database"]["dbname"], config["database"]["port"])

	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
	}

	// Auto-migrate models
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
