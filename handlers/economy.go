package handlers

import (
	"sbox-backend/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func registerEconomyHandlers(f *fiber.App) {
	g := f.Group("/economy")

	g.Post("/pay", handlePay)
}

func handlePay(c *fiber.Ctx) error {
	id := c.Get("Steam-ID64")

	recipient := c.Query("id")
	amount, err := strconv.Atoi(c.Query("amount"))
	if err != nil {
		return err
	}

	money, err := db.GetMoney(id)
	if err != nil {
		return err
	}

	if (money - amount) <= 0 {
		return c.SendString("no money waaaa")
	}

	err = db.AddMoney(recipient, amount)
	if err != nil {
		return err
	}

	err = db.SubtractMoney(id, amount)
	if err != nil {
		return err
	}

	err = db.LogPayment(id, recipient, amount)
	if err != nil {
		return err
	}

	return nil
}
