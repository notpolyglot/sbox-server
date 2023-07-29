package main

import (
	"log"
	"sbox-backend/db"
	"sbox-backend/structs"

	"github.com/gofiber/fiber/v2"
)

func (s Server) setupHandlers() {
	f := fiber.New()

	f.Get("/login", s.handleLogin)

	log.Fatal(f.Listen(":3000"))

}

func (s *Server) handleLogin(c *fiber.Ctx) error {
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
