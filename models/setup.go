package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	fmt.Println("Connecting to database...")
	database, err := gorm.Open("sqlite3", "town.db")

	fmt.Println("Connected to database.")
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Person{})
	database.AutoMigrate(&Building{})
	database.AutoMigrate(&BuildingType{})
	database.AutoMigrate(&Job{})

	DB = database
}
