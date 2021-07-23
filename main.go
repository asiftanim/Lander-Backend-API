package main

import (
	"lander/configs"
	"lander/database"
	"lander/routes"

	"log"
)

func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Connot load config", err)
	}

	database.Connect(config.DBSource)
	server := routes.Init()

	log.Fatal(server.Run(config.SERVER))
}
