package main

import (
    "log"
    "time"
    "github.com/gofiber/fiber/v3"
)

type Response struct {
  Message   string `json:"message"`
  Timestamp int64  `json:"timestamp"`
}


func main() {
    app := fiber.New()

    app.Get("/", func (c fiber.Ctx) error {
        return c.JSON(Response{
			Message:   "My name is Rainier Ring",
			Timestamp: time.Now().Unix(),
		})
    })

    log.Fatal(app.Listen(":3000"))
}
