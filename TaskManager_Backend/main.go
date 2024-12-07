package main

import (
	"taskmanager/config"
	"taskmanager/routes"
)

func main() {
	config.ConnectDB()

	router := routes.SetupRouter()
	router.Run(":8080")
}
