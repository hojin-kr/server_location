package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hojin-kr/location"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/location/:uuid", func(c *fiber.Ctx) error {
		location := location.Location{}
		location.UUID = c.Params("uuid")
		location.GetLocation()
		return c.SendString(location.Location)
	})

	app.Post("/location", func(c *fiber.Ctx) error {
		location := location.Location{}
		if err := c.BodyParser(&location); err != nil {
			return err
		}
		location.SetLocation()
		return c.SendString(location.Location)
	})

	app.Listen(":3000")
}
