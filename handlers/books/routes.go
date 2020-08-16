package books

import "github.com/raene/fiberWeb/config"

/*
Book type holds values needed for the app handlers to function,
in this case Router and Db values
*/
type Book struct {
	Config *config.App
}

//SetupRoutes sets up all routes
func (a *Book) SetupRoutes() {
	books := a.Config.Router.Group("/books")
	books.Get("/", a.helloWorld)
	books.Post("/", a.createBook)
}
