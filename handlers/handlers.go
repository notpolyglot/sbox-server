package handlers

import "github.com/gofiber/fiber/v2"

func RegisterHandlers(f *fiber.App) {
	registerUserHandlers(f)
	registerEconomyHandlers(f)
}
