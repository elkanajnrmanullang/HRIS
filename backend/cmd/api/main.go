package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"celestara.com/hris-api/internal/config"
	"celestara.com/hris-api/internal/infrastructure/database"
	"celestara.com/hris-api/internal/shared/response"

	authHandler "celestara.com/hris-api/internal/module/auth/handler"
	authRepo "celestara.com/hris-api/internal/module/auth/repository"
	authService "celestara.com/hris-api/internal/module/auth/service"

	tenantHandler "celestara.com/hris-api/internal/module/tenant/handler"
	tenantRepo "celestara.com/hris-api/internal/module/tenant/repository"
	tenantService "celestara.com/hris-api/internal/module/tenant/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("..")
	if err != nil {
		log.Fatalf("Config Error: %v", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Database Error: %v", err)
	}

	// RunDevMigration
	database.RunDevMigration(db, cfg.App.Env)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("SQL DB Error: %v", err)
	}
	defer sqlDB.Close()

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Dependency Injection Setup
	companyRepo := tenantRepo.NewCompanyRepository(db)
	companyService := tenantService.NewCompanyService(companyRepo)
	companyHandler := tenantHandler.NewCompanyHandler(companyService)

	userRepo := authRepo.NewUserRepository(db)
	authSvc := authService.NewAuthService(userRepo)
	authHdlr := authHandler.NewAuthHandler(authSvc)

	// HealthCheck
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Success("Server is healthy", nil))
	})

	// API Routes V1
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHdlr.Register)
			auth.POST("/login", authHdlr.Login)
		}

		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/companies", middleware.RoleMiddleware("tenant_admin"), companyHandler.Create) 
		}
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: router,
	}

	// ServerGoroutine
	go func() {
		log.Printf("Server running on port %d", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server Listen Error: %v", err)
		}
	}()

	// GracefulShutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Error: %v", err)
	}
}