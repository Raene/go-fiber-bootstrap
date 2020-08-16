package config

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router fiber.Router
	Db     *gorm.DB
}

func AppInit(router fiber.Router, c chan *gorm.DB) *App {
	app := App{}
	app.Router = router
	app.Db = <-c
	return &app
}
