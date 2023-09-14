package main

import (
	"chenwoo-simple-invoices/domain"
	"chenwoo-simple-invoices/handler"
	"chenwoo-simple-invoices/mongodb"
	"chenwoo-simple-invoices/repository"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ir          domain.Invoices
	ih          handler.InvoicesHandler
	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)

func main() {
	ctx = context.Background()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mongoclient, err := mongodb.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	db := mongoclient.Database("invoices")

	ir = repository.NewInvoices(db)
	ih = handler.NewInvoicesHandler(ir)
	server = gin.Default()

	c, _ := context.WithTimeout(ctx, 10*time.Second)

	defer mongoclient.Disconnect(c)
	err = mongoclient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	basepath := server.Group("/v1")

	ih.InvoicesRoutes(basepath)
	server.Run(":3000")
}
