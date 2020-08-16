package products

import "github.com/gofiber/fiber"

/*
App type holds values needed for the app handlers to function
in this case Router and Db values
*/
type App struct {
	Router fiber.Router
}

//SetupRoutes sets up all routes
func (a *App) SetupRoutes() {
	product := a.Router.Group("/product")
	product.Get("/", a.helloWorld)
}
