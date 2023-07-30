package handlers

import (
	"sbox-backend/db"
	"sbox-backend/structs"

	"github.com/gofiber/fiber/v2"
)

func registerUserHandlers(f *fiber.App) {
	g := f.Group("/user")

	g.Post("/login", handleLogin)
}

func handleLogin(c *fiber.Ctx) error {
	id := c.Get("Steam-ID64")

	_, err := db.GetUser(id)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err := db.CreateUser(id)
			if err != nil {
				return err
			}
		}
	}

	money, err := db.GetMoney(id)
	if err != nil {
		return err
	}

	return c.JSON(structs.LoginResponse{
		Money: money,
	})
}
