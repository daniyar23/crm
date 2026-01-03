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
	// test
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	delivery.InitRoutes(router)

	// User
	userRepo := db.NewUserPostgresRepository(dbConn)
	userService := services.NewUserService(userRepo)
	userHandler := delivery.NewUserHandler(userService)
	api := router.Group("/api")
	userHandler.RegisterRoutes(api)

	// Company
	companyRepo := db.NewCompanyPostgresRepository(dbConn)
	companyService := services.NewCompanyService(companyRepo)
	companyHandler := delivery.NewCompanyHandler(companyService)
	companyHandler.RegisterRoutes(api)

	// HTML страницы
	router.LoadHTMLFiles("static/index.html")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Запуск сервера
	if err := router.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}

}
