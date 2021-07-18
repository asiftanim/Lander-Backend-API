package database

import (
	"gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect(DBSource string){
	connection, err := gorm.Open(mysql.Open(DBSource), &gorm.Config{})

  if err != nil{
    panic("Cannot connect to database");
  }

  DB = connection
}