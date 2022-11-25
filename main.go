package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("templates", ".html"),
	})

	app.Use(func(c *fiber.Ctx) error {
		if err := c.Bind(fiber.Map{"Title": "default"}); err != nil {
			return fmt.Errorf("failed to bind default vars: %w", err)
		}
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"Title": "Home"})
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
