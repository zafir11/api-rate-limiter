package main

import (
	"rate-limiter/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	tokenBucket := middleware.NewTokenbucket(10, 2)

	// Middleware to perform rate limiting
	app.Use(func(c *fiber.Ctx) error {
		if tokenBucket.Request(1) { // Specify the number of tokens needed for each request
			return c.Next()
		}
		return c.SendStatus(fiber.StatusTooManyRequests)
	})
	app.Get("/limited", func(c *fiber.Ctx) error {
		return c.SendString("you are Limited")
	})
	app.Get("/unlimited", func(c *fiber.Ctx) error {
		return c.SendString("you are Unlimited")
	})
	app.Listen(":8080")
}
