package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/raene/fiberWeb/database"
	"github.com/raene/fiberWeb/handlers/books"
	"github.com/raene/fiberWeb/handlers/products"
)

func main() {
	database.InitDatabase()

	db := database.DBConn

	a := &products.App{}
	b := &books.App{}
	app := fiber.New()
	api := app.Group("/api", logger.New())
	a.Router = api
	b.Router = api

	b.SetupRoutes()
	a.SetupRoutes()
	app.Listen(3000)

	defer db.Close()
}
