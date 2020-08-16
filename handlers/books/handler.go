package books

import (
	"github.com/gofiber/fiber"
	"github.com/raene/fiberWeb/models"
)

func (a *App) helloWorld(c *fiber.Ctx) {
	var books []models.Book
	a.Db.Find(&books)
	c.JSON(books)
}

func (a *App) createBook(c *fiber.Ctx) {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	a.Db.Create(&book)
	c.JSON(book)
}
