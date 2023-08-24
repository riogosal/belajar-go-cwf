package main

import (
	"eser-http/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routes.SetupRoutes()
}
