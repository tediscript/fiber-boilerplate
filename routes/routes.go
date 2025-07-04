package routes

import (
	"fiber-boilerplate/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// SetupRoutes sets up the routes for the Fiber app
func SetupRoutes(app *fiber.App) {
	authHandler := handlers.NewAuthHandler()
	homeHandler := handlers.NewHomeHandler()
	dashboardHandler := handlers.NewDashboardHandler()

	// Define routes
	app.Get("/", homeHandler.Home)
	app.Get("/auth/login", authHandler.Login)
	app.Get("/auth/register", authHandler.Register)
	app.Get("/auth/logout", authHandler.Logout)
	app.Get("/dashboard", dashboardHandler.Dashboard)

	// New login page inspired by template
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("pages/login", fiber.Map{})
	})

	// New register page inspired by template
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("pages/register", fiber.Map{})
	})

	// Add health check endpoints (/livez and /readyz)
	app.Use(healthcheck.New())

	// Add metrics dashboard endpoint (/metrics)
	app.Get("/metrics", monitor.New())
}
