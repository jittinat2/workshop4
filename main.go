package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Basic routes (Hello World example)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	port := 3000
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on %s", addr)
	log.Fatal(app.Listen(addr))
}

// ...existing code... (handlers are in users.go)
