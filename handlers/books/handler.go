package books

import (
	"github.com/gofiber/fiber"
	"github.com/raene/fiberWeb/models"
)

func (a *Book) helloWorld(c *fiber.Ctx) {
	var books []models.Book
	a.Config.Db.Find(&books)
	c.JSON(books)
}

func (a *Book) createBook(c *fiber.Ctx) {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	a.Config.Db.Create(&book)
	c.JSON(book)
}
