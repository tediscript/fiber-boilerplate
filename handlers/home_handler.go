package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	// Add any dependencies here
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(c *fiber.Ctx) error {
	return c.Render("pages/home", fiber.Map{
		"Title": "Fiber Boilerplate",
		"Year":  time.Now().Year(),
	})
}
