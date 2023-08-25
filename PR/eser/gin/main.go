package main

import (
	"eser-gin/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routes.SetupRoutes()

}
