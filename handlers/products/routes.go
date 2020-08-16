package products

import "github.com/raene/fiberWeb/config"

/*
App type holds values needed for the app handlers to function
in this case Router and Db values
*/
type Product struct {
	Config *config.App
}

//SetupRoutes sets up all routes
func (a *Product) SetupRoutes() {
	product := a.Config.Router.Group("/product")
	product.Get("/", a.helloWorld)
}
