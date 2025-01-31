package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/renatorrocha/minipay/pkg/database"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	log.Fatal(app.Listen(":3001"))
}
