package main

import (
	"rate-limiter/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fixedWindowLimiter := middleware.NewSlidingWindow(time.Minute, 5)

	// Middleware to perform rate limiting
	app.Use(func(c *fiber.Ctx) error {
		if fixedWindowLimiter.Allow() {
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
