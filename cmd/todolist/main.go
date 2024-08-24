package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to Home")
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Println(err)
	}
}
