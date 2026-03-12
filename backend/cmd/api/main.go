package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/config"
	"github.com/username/multi-inventory-manager/internal/domain"
	"github.com/username/multi-inventory-manager/internal/repository"
	"github.com/username/multi-inventory-manager/internal/usecase"
	"github.com/username/multi-inventory-manager/internal/delivery/http/middleware"
	deliveryHttp "github.com/username/multi-inventory-manager/internal/delivery/http"
	"github.com/username/multi-inventory-manager/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Connect to Postgres
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Could not connect to Postgres: %v", err)
	}

	// Auto Migration
	err = db.AutoMigrate(
		&domain.Tenant{}, 
		&domain.User{}, 
		&domain.Item{}, 
		&domain.Product{}, 
		&domain.Recipe{}, 
		&domain.Table{}, 
		&domain.Order{}, 
		&domain.OrderItem{}, 
		&domain.PreparationLog{},
	)
	if err != nil {
		log.Fatalf("Could not auto migrate: %v", err)
	}

	// Connect to Redis
	rdb, err := database.ConnectRedis(cfg)
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	_ = rdb

	// Initialize Repository
	authRepo := repository.NewAuthRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	tenantRepo := repository.NewTenantRepository(db)

	// Initialize Usecase
	authUsecase := usecase.NewAuthUsecase(authRepo, cfg.JWTSecret)
	inventoryUsecase := usecase.NewInventoryUsecase(inventoryRepo)
	productUsecase := usecase.NewProductUsecase(productRepo, inventoryRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, inventoryRepo)
	tenantUsecase := usecase.NewTenantUsecase(tenantRepo)
	reportUsecase := usecase.NewReportUsecase(db)

	// Initialize Handler
	authHandler := deliveryHttp.NewAuthHandler(authUsecase)
	inventoryHandler := deliveryHttp.NewInventoryHandler(inventoryUsecase)
	productHandler := deliveryHttp.NewProductHandler(productUsecase)
	orderHandler := deliveryHttp.NewOrderHandler(orderUsecase)
	tenantHandler := deliveryHttp.NewTenantHandler(tenantUsecase)
	dashboardHandler := deliveryHttp.NewDashboardHandler(reportUsecase)

	// Setup Gin
	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	// CORS Middleware
	r.Use(middleware.CORSMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	// Auth Routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Protected Routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		api.GET("/me", func(c *gin.Context) {
			userID := c.MustGet("user_id")
			tenantID := c.MustGet("tenant_id")
			role := c.MustGet("role")
			c.JSON(200, gin.H{
				"user_id":   userID,
				"tenant_id": tenantID,
				"role":      role,
			})
		})

		// Inventory Routes
		inventory := api.Group("/inventory")
		{
			inventory.POST("/items", inventoryHandler.CreateItem)
			inventory.GET("/items", inventoryHandler.ListItems)
			inventory.GET("/items/barcode/:barcode", inventoryHandler.GetByBarcode)
			inventory.PATCH("/items/:id/stock", inventoryHandler.UpdateStock)
		}

		// Product & Recipe Routes
		products := api.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)
			products.GET("", productHandler.ListProducts)
			products.POST("/:id/recipe", productHandler.SetRecipe)
			products.POST("/:id/prepare", productHandler.Prepare)
		}

		api.GET("/preparations", productHandler.ListPreparations)

		// Tenant Routes
		api.GET("/tenant", tenantHandler.GetTenant)
		api.PATCH("/tenant", tenantHandler.UpdateTenant)

		// Table & Order Routes
		tables := api.Group("/tables")
		{
			tables.POST("", orderHandler.CreateTable)
			tables.GET("", orderHandler.ListTables)
		}

		orders := api.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.GET("", orderHandler.ListOrders)
			orders.POST("/:id/complete", orderHandler.CompleteOrder)
		}

		// Dashboard Routes
		dashboard := api.Group("/dashboard")
		{
			dashboard.GET("/summary", dashboardHandler.GetSummary)
		}
	}

	log.Printf("Starting server on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
