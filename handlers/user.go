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
	req := new(structs.LoginReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user, err := db.GetUser(req.Id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
