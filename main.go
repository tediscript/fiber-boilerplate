package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	"fiber-boilerplate/handlers"
)

// Helper function to get environment variable with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	// Get environment variables with fallbacks
	port := getEnv("PORT", "3000")
	appName := getEnv("APP_NAME", "Fiber Boilerplate")
	shutdownTimeoutStr := getEnv("SHUTDOWN_TIMEOUT", "5")

	// Parse shutdown timeout
	shutdownTimeout, err := strconv.Atoi(shutdownTimeoutStr)
	if err != nil {
		shutdownTimeout = 5 // Default to 5 seconds if parsing fails
	}

	// Initialize template engine
	views := html.New("./views", ".html")

	// Create a new Fiber app with template engine
	app := fiber.New(fiber.Config{
		AppName: appName,
		Views:   views,
	})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler()
	homeHandler := handlers.NewHomeHandler()
	dashboardHandler := handlers.NewDashboardHandler()

	// Define routes
	app.Get("/", homeHandler.Home)
	app.Get("/auth/login", authHandler.Login)
	app.Get("/auth/register", authHandler.Register)
	app.Get("/auth/logout", authHandler.Logout)
	app.Get("/dashboard", dashboardHandler.Dashboard)

	// Define a route for the root path

	// Add health check endpoints (/livez and /readyz)
	app.Use(healthcheck.New())

	// Add metrics dashboard endpoint (/metrics)
	app.Get("/metrics", monitor.New())

	// Setup graceful shutdown
	// Create channel for shutdown signals
	shutdownChan := make(chan os.Signal, 1)
	// Notify the channel for system signals
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Server %s starting on http://localhost:%s\n", appName, port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-shutdownChan
	log.Println("Shutdown signal received")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeout)*time.Second)
	defer cancel()

	// Shutdown the server with timeout
	log.Printf("Shutting down server (timeout: %d seconds)...\n", shutdownTimeout)
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
