package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(DBSource string) {
	connection, err := gorm.Open(mysql.Open(DBSource), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	DB = connection
}
