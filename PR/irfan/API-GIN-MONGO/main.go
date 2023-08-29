package main

import (
	"app-gin-mongo/db"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("DB_NAME"))
	client := db.MgoConn()
	db.MgoCollection()
	defer client.Disconnect(context.TODO())
}
