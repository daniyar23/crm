package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/daniyar23/crm/internal/core/config"
	delivery "github.com/daniyar23/crm/internal/feature/feature1/delivery/http-grps"
	"github.com/daniyar23/crm/internal/feature/feature1/infrastructure/db"
	"github.com/daniyar23/crm/internal/feature/feature1/services"
	"github.com/daniyar23/crm/internal/feature/feature1/usecase"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ---------- config ----------
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// ---------- db ----------
	dbConn, err := db.NewPostgres(cfg.DB.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// ---------- repositories ----------
	userRepo := db.NewUserPostgresRepository(dbConn)
	companyRepo := db.NewCompanyPostgresRepository(dbConn)

	// ---------- services (business logic) ----------
	userService := services.NewUserService(userRepo)
	companyService := services.NewCompanyService(companyRepo)

	// ---------- event bus ----------
	eventBus := usecase.NewEventBus(100)

	// ---------- usecases ----------
	userUC := usecase.NewUserUseCase(userService, eventBus)
	companyUC := usecase.NewCompanyUseCase(companyService)

	// ---------- async listeners ----------
	usecase.StartUserDeletedListener(
		ctx,
		eventBus,
		companyService, // service ОК — это background business flow
	)

	// ---------- http ----------
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	// users
	userHandler := delivery.NewUserHandler(userUC)
	userHandler.RegisterRoutes(api)

	// companies
	companyHandler := delivery.NewCompanyHandler(companyUC)
	companyHandler.RegisterRoutes(api)

	// ---------- static / html (опционально) ----------
	router.LoadHTMLFiles("static/index.html")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// ---------- start ----------
	if err := router.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
