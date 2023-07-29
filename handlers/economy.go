package handlers

import "github.com/gofiber/fiber/v2"

func registerEconomyHandlers(f *fiber.App) {
	g := f.Group("/economy")

	g.Post("/pay", handlePay)
}

func handlePay(c *fiber.Ctx) error {
	return nil
}
