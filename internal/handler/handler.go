package handler

import (
	"ex_proj_go/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/registration")
		auth.POST("/login")
		auth.POST("/refresh")
		auth.DELETE("/logout")
		auth.PATCH("/replace_password")

	}

	users := router.Group("/users")
	{
		users.GET("/user")
		users.GET("/user:id")
		users.PATCH("/users")
		users.PATCH("/location")

	}

	return router
}
