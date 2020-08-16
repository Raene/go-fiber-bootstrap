package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/jinzhu/gorm"
	Config "github.com/raene/fiberWeb/config"
	"github.com/raene/fiberWeb/database"
	"github.com/raene/fiberWeb/handlers/books"
	"github.com/raene/fiberWeb/handlers/products"
)

type Routes interface {
	SetupRoutes()
}

func spawnRoutes(m chan string, r ...Routes) {
	for _, v := range r {
		v.SetupRoutes()
	}
	m <- "Routes Setup"
}

func main() {
	var c chan *gorm.DB = make(chan *gorm.DB)
	var m chan string = make(chan string)
	go database.InitDatabase(c)

	app := fiber.New()
	api := app.Group("/api", logger.New())
	config := Config.AppInit(api, c)

	b := &books.Book{Config: config}
	p := &products.Product{Config: config}

	go spawnRoutes(m, b, p)

	fmt.Println(<-m)
	app.Listen(3000)

}
