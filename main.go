package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Hello World path
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	addr := ":3000"
	log.Printf("Starting server on %s", addr)
	log.Fatal(app.Listen(addr))
}
