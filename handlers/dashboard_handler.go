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
	// Mock stats data
	stats := []struct {
		Title string
		Value string
	}{
		{"Total Sales", "$1,234,567"},
		{"New Users", "89"},
		{"Pending Orders", "12"},
	}

	// Mock orders data
	orders := []struct {
		ID       int
		Customer string
		Status   string
		Total    string
	}{
		{101, "Jane Doe", "Shipped", "$250.00"},
		{102, "John Smith", "Processing", "$145.50"},
		{103, "Alice Johnson", "Pending", "$89.99"},
		{104, "Bob Brown", "Completed", "$320.50"},
		{105, "Charlie Davis", "Cancelled", "$75.25"},
	}

	return c.Render("pages/dashboard", fiber.Map{
		"Stats":  stats,
		"Orders": orders,
	}, "layouts/admin")
}
