package books

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

/*
App type holds values needed for the app handlers to function
in this case Router and Db values
*/
type App struct {
	Router fiber.Router
	Db     *gorm.DB
}

//SetupRoutes sets up all routes
func (a *App) SetupRoutes() {
	books := a.Router.Group("/books")
	books.Get("/", a.helloWorld)
	books.Post("/", a.createBook)
}
