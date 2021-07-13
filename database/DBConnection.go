package database

import (
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect(){
	connection, err := gorm.Open(mysql.Open("root:root@/lander"), &gorm.Config{})

  if err != nil{
    panic("Cannot connect to database");
  }

  DB = connection
}