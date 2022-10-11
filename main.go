package main

import (
	"assigment-2/database"
	"assigment-2/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
