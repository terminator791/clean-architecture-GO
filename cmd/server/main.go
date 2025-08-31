package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/terminator791/clean-architecture-GO/config"
	"github.com/terminator791/clean-architecture-GO/internal/adapter/controller"
	"github.com/terminator791/clean-architecture-GO/internal/adapter/repository"
	"github.com/terminator791/clean-architecture-GO/internal/infrastructure/database"
	"github.com/terminator791/clean-architecture-GO/internal/infrastructure/router"
	"github.com/terminator791/clean-architecture-GO/internal/usecase"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := database.NewPostgresConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Health check database
	if err := database.HealthCheck(db); err != nil {
		log.Fatalf("Database health check failed: %v", err)
	}
	log.Println("Database connection established")

	// Initialize repositories
	productRepo := repository.NewPostgresProductRepository(db)

	// Initialize use cases
	productUseCase := usecase.NewProductUseCase(productRepo)

	// Initialize controllers
	productController := controller.NewProductController(productUseCase)

	// Initialize router
	httpRouter := router.NewRouter(productController)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      httpRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %d", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server shutting down...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}