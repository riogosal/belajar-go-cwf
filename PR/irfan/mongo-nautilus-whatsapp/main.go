package main

import (
	"app-api-natulius/connection"
	"app-api-natulius/domain"
	"app-api-natulius/repositories"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	rs          domain.RawMaterialsService
	ctx         context.Context
	rawMaterial *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoclient = connection.Connecting(ctx, mongoclient, err)
	rawMaterial = mongoclient.Database("natulius").Collection("raw_material_lots")
	rs = repositories.NewRawMaterialService(rawMaterial, ctx)
}

func main() {

	c, _ := context.WithTimeout(ctx, 10*time.Second)

	defer mongoclient.Disconnect(c)
	err = mongoclient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	id := "6417b194954fea8407a5aadc"
	nautilus, err := rs.GetData(id)
	if err != nil {
		panic(err)
	}

	fmt.Println(nautilus.CetakStrRwaMatetirialLot())

}
