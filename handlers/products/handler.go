package products

import "github.com/gofiber/fiber"

func (a *App) helloWorld(c *fiber.Ctx) {
	c.Send("Hello, Wor")
}
