package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func registerXPHandler(f *fiber.App) {
	g := f.Group("/xp")

	g.Post("/increase", handleLogin)
}

// func handleIncreaseXP(c *fiber.Ctx) error {
// 	id := c.Get("Steam-ID64")

// }
