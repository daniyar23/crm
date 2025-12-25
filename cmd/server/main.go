package main

import (
	"github.com/gin-gonic/gin"

	"github.com/daniyar23/crm/internal/handlers"
	"github.com/daniyar23/crm/internal/repository"
	"github.com/daniyar23/crm/internal/services"
)

func main() {
	r := gin.Default()

	userRepo := repository.NewUserMemoryRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	api := r.Group("/api")
	userHandler.RegisterRoutes(api)

	r.Run(":8080")
}
