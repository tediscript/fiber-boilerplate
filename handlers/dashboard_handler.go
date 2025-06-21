package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	// Add any dependencies here
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) Dashboard(c *fiber.Ctx) error {
	// Implement dashboard logic here
	return c.SendString("Dashboard handler")
}
