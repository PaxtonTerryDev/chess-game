package routes

import (
	"github.com/gin-gonic/gin"

	"chess-game/server/internal/api/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	
	api.GET("/health", handlers.HealthCheck)
	
	// Future routes will be added here:
	// auth := api.Group("/auth")
	// game := api.Group("/game")
	// shop := api.Group("/shop")
	// player := api.Group("/player")
}