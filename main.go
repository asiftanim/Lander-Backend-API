package main

import (
	"fmt"
	"lander/configs"
	"lander/database"
	"lander/routes"
	"log"
	"os"
)

func main() {

	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Connot load config", err)
	}

	InitAssetFolder(config.IMAGE_PATH)
	database.Connect(config.DBSource)
	server := routes.Init()

	log.Fatal(server.Run(config.SERVER))
}

func InitAssetFolder(image_path string) {
	_, err := os.Stat(image_path)
	if os.IsNotExist(err) {
		fmt.Println("path does not exist")
		os.MkdirAll(image_path, os.ModePerm)
	}
}
