package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	router := fiber.New()

	router.Use(cors.New())

	// GET /api/register
	router.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!!!") // => ✋ register
	})

	router.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(router.Listen(":4001"))
}
