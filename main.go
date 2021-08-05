package main

import (
	"fmt"
	"lander/configs"
	"lander/database"
	"lander/routes"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// getting env variables IMAGE_PATH
	image_path := os.Getenv("IMAGE_PATH")
	folder_info, err := os.Stat(image_path)
	if os.IsNotExist(err) {
		fmt.Println("path does not exist")
		os.MkdirAll(image_path, os.ModePerm)
	}

	if !os.IsNotExist(err) {
		fmt.Println(folder_info)
	}

	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Connot load config", err)
	}

	database.Connect(config.DBSource)
	server := routes.Init()

	log.Fatal(server.Run(config.SERVER))
}
