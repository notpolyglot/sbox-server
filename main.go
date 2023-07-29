package main

import (
	"log"
	"sbox-backend/db"
	"sbox-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	f := fiber.New()

	handlers.RegisterHandlers(f)

	log.Fatal(f.Listen(":3000"))
}
