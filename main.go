package main

import (
  "lander/database"
  "lander/routes"
  "lander/configs"
  
  "log"
  
)

func main() {
  config, err := configs.LoadConfig(".")
  
  if err != nil{
    log.Fatal("Connot load config", err)
  }

  database.Connect(config.DBSource)
  server := routes.Init()

  log.Fatal(server.Run(config.SERVER))
}