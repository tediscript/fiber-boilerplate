package main

import (
	// Standard library
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// Project packages
	"fiber-boilerplate/configs"
	"fiber-boilerplate/routes"

	// External packages
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Configuration Initialization
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	// Initialize configuration
	cfg, err := configs.InitConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Template Engine Setup
	views := html.New("./views", ".html")

	// Fiber Application Creation
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
		Views:   views,
	})

	// Static Files Configuration
	app.Static("/static", "./static")

	// Route Registration
	routes.SetupRoutes(app)

	// Graceful Shutdown Configuration
	// Create channel for shutdown signals
	shutdownChan := make(chan os.Signal, 1)
	// Notify the channel for system signals
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	// Server Startup
	go func() {
		log.Printf("Server %s starting on http://localhost:%s\n", cfg.AppName, cfg.Port)
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Shutdown Handling
	// Wait for shutdown signal
	<-shutdownChan
	log.Println("Shutdown signal received")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
	defer cancel()

	// Shutdown the server with timeout
	log.Printf("Shutting down server (timeout: %d seconds)...\n", cfg.ShutdownTimeout)
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
