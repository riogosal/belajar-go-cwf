package main

import (
	"app-api-natulius/connection"
	"app-api-natulius/domain"
	"app-api-natulius/models"
	"app-api-natulius/repositories"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx         context.Context
	db          *mongo.Database
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.Background()
	mongoclient = connection.Connecting(ctx, mongoclient, err)
	db = mongoclient.Database("natulius")
}

func main() {

	// Repositories
	var (
		rs  domain.RawMaterialsService
		psr domain.ProductService
		rm  domain.RefinedMaterialServices
	)

	rs = repositories.NewRawMaterialService(db)
	psr = repositories.NewProductServiceImpl(db)
	rm = repositories.NewRefinedMaterialLots(db)

	c, _ := context.WithTimeout(ctx, 10*time.Second)

	defer mongoclient.Disconnect(c)
	err = mongoclient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// Get Data By ID
	// id := "6417b194954fea8407a5aadc"
	// nautilus, err := rs.GetData(ctx, id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(nautilus.CetakStrRwaMatetirialLot())

	// Get Data By Date
	timeZone, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		panic(err)
	}

	startDate := time.Date(2023, time.April, 15, 0, 0, 0, 0, timeZone)
	endDate := time.Date(2023, time.April, 15, 23, 59, 59, 9999999999, timeZone)

	IntStartDate := startDate.UnixMilli()
	IntEndDate := endDate.UnixMilli()

	// Data Raw Materials
	dataMaterials, err := rs.GetDataByDate(ctx, IntStartDate, IntEndDate)
	if err != nil {
		panic(err)
	}

	for _, rawMaterialFilterDate := range dataMaterials {
		fmt.Println(rawMaterialFilterDate.CetakStrRwaMatetirialLot())
	}

	// Data Refined Material
	refined, err := rm.GetDataByDate(ctx, IntStartDate, IntEndDate)
	if err != nil {
		panic(err)
	}

	for _, rml := range refined {
		fmt.Println(rml.PrintDataRefinedMaterial())
	}

	// Data Packing
	productDate, err := psr.GetDataByDate(ctx, IntStartDate, IntEndDate)
	if err != nil {
		panic(err)
	}
	models.DataProductsdate(productDate)
	// models.DataGroup()

}
