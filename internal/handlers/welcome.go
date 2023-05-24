package handlers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	// return c.SendString("Welcome to Fiber!")
	return c.Render("welcome", nil, "layouts/main")
}
