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
	r := gin.Default()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.NewPostgres(cfg.DB.DSN())
	if err != nil {
		log.Fatal(err)
	}

	userRepo := db.NewUserPostgresRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := delivery.NewUserHandler(userService)

	api := r.Group("/api")
	userHandler.RegisterRoutes(api)

	r.Run(":8080")
}
