package main

import (
  "lander/database"
  "lander/routes"
   "log"
)

func main() {

  database.Connect()
  server := routes.Init()

  log.Fatal(server.Run(":8080"))
}