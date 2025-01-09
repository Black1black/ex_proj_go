package main

import (
	"ex_proj_go/internal/db"
	"ex_proj_go/internal/handlers"
	"ex_proj_go/internal/services"
	"ex_proj_go/pkg/logger"
	"log"

	//_ "ex_proj_go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ExProjectGo API
// @version 1.0
// @description This is a sample API with Gin, GORM, Swagger, and Worker Pool.

// @host localhost:8080
// @BasePath /api/v1

func main() {
	logger.InitLogger()

	// Initialize database
	db.InitDB()

	// Initialize Gin router
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.Group("/api/v1")
	{
		handlers.RegisterUserRoutes(api)
		handlers.RegisterProductRoutes(api)
	}

	// Start worker pool
	services.InitWorkerPool(5)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
