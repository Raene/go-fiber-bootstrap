package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/jinzhu/gorm"
	"github.com/raene/fiberWeb/database"
	"github.com/raene/fiberWeb/handlers/books"
	"github.com/raene/fiberWeb/handlers/products"
)

func main() {
	var c chan *gorm.DB = make(chan *gorm.DB)
	go database.InitDatabase(c)

	db := <-c

	a := &products.App{}
	b := &books.App{}
	app := fiber.New()
	api := app.Group("/api", logger.New())
	a.Router = api
	b.Router = api
	b.Db = db

	go b.SetupRoutes()
	go a.SetupRoutes()
	app.Listen(3000)

	defer db.Close()
}
