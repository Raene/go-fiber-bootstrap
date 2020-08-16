package products

import "github.com/gofiber/fiber"

func (a *Product) helloWorld(c *fiber.Ctx) {
	c.Send("Hello, Wor")
}
