package main

import (
	"log"
	"sbox-backend/db"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := Server{
		State: GameState{},
	}

	db.Connect()

	s.setupHandlers()

}
