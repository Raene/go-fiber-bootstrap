package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/raene/fiberWeb/models"
)

func InitDatabase(c chan *gorm.DB) {
	var err error
	DBConn, err := gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Connection Opened to Database")
	c <- DBConn
}
