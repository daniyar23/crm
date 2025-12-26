package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/daniyar23/crm/internal/config"
	delivery "github.com/daniyar23/crm/internal/delivery/http"
	"github.com/daniyar23/crm/internal/infrastructure/db"
	"github.com/daniyar23/crm/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	dbConn, err := db.NewPostgres(cfg.DB.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	userRepo := db.NewUserPostgresRepository(dbConn)
	userService := services.NewUserService(userRepo)
	userHandler := delivery.NewUserHandler(userService)

	api := router.Group("/api")
	userHandler.RegisterRoutes(api)

	if err := router.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
