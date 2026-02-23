package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c fiber.Ctx) error {
        return c.JSON(fiber.Map{
		"message" : "My name is Rainier Ring",
		"timestamp" : "_____",
	})
    })

    log.Fatal(app.Listen(":3000"))
}
