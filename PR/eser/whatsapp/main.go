package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"whatsapp-nautilus/domain"
	"whatsapp-nautilus/mongodb"
	"whatsapp-nautilus/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	r           domain.RawMaterialLotRepository
	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoclient, err = mongodb.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	rawMaterial := mongoclient.Database("nautilus").Collection("raw_material_lots")
	r = repository.NewRawMaterialLot(rawMaterial, ctx)
}

func main() {
	c, _ := context.WithTimeout(ctx, 10*time.Second)
	defer mongoclient.Disconnect(c)
	err = mongoclient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	id := "641cfe6c88032b19610a6499" //  tolong kenapa ini eroro terus
	mydata, err := r.GetData(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mydata.CetakStrRawmaterialLot())
}
