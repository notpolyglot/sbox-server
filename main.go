package main

import (
	"log"
	"sbox-backend/db"
	"sbox-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func handleFiberError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	f := fiber.New(fiber.Config{
		ErrorHandler: handleFiberError,
	})
	f.Use(logger.New())
	// Logging Request ID

	handlers.RegisterHandlers(f)

	log.Fatal(f.Listen(":80"))
}
