package main

import (
	"app-api-natulius/connection"
	"app-api-natulius/domain"
	"app-api-natulius/handler"
	"app-api-natulius/repositories"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	db          *mongo.Database
	mongoclient *mongo.Client
	err         error
)

func init() {
	fmt.Println("Selamat Sore Chen Woo Fishery (Makassar) ini rangkmuman processing tanggal 2023-06-23")
	ctx = context.Background()
	mongoclient = connection.Connecting(ctx, mongoclient, err)
	db = mongoclient.Database("natulius")
	server = gin.Default()
}

func main() {
	defer fmt.Println("\nSekedar informasinya data produksi untuk hari ini, Terima Kasih")
	// Repositories
	var (
		ri domain.InvoiceService
		hr handler.InvoiceHandler
	// rs  domain.RawMaterialsService
	// psr domain.ProductService
	// rm  domain.RefinedMaterialServices
	// rp  domain.PackingService
	)

	// rs = repositories.NewRawMaterialService(db)
	// psr = repositories.NewProductServiceImpl(db)
	// rm = repositories.NewRefinedMaterialLots(db)
	// rp = repositories.NewPackingServiceImpl(db)
	ri = repositories.NewInvoiceSerciceImpl(db)
	hr = handler.NewInvoiceHandler(ri)

	c, _ := context.WithTimeout(ctx, 10*time.Second)

	defer mongoclient.Disconnect(c)
	err = mongoclient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	basepath := server.Group("/v1")
	hr.InvoiceRoutes(basepath)
	server.Run(":3001")

	// Get Data By ID
	// id := "6417b194954fea8407a5aadc"
	// nautilus, err := rs.GetData(ctx, id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(nautilus.CetakStrRwaMatetirialLot())

	// Get Data By Date
	// timeZone, err := time.LoadLocation("Asia/Makassar")
	// if err != nil {
	// 	panic(err)
	// }

	// startDate := time.Date(2023, time.July, 15, 0, 0, 0, 0, timeZone)
	// endDate := time.Date(2023, time.July, 15, 23, 59, 59, 9999999999, timeZone)

	// IntStartDate := startDate.UnixMilli()
	// IntEndDate := endDate.UnixMilli()

	// // Data Raw Materials
	// dataMaterials, err := rs.GetDataByDate(ctx, IntStartDate, IntEndDate)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, rawMaterialFilterDate := range dataMaterials {
	// 	fmt.Println(rawMaterialFilterDate.CetakStrRwaMatetirialLot())
	// }

	// // Data Refined Material
	// refined, err := rm.GetDataByDate(ctx, IntStartDate, IntEndDate)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, rml := range refined {
	// 	fmt.Println(rml.PrintDataRefinedMaterial())
	// }

	// // Data Product/Retouching
	// productDate, err := psr.GetDataByDate(ctx, IntStartDate, IntEndDate)
	// if err != nil {
	// 	panic(err)
	// }
	// models.DataProductsdate(productDate)
	// // models.DataGroup()

	// // Data Packings
	// dataPacking, err := rp.GetDataByDate(ctx, IntStartDate, IntEndDate)
	// if err != nil {
	// 	panic(err)
	// }

	// models.PrintDataPackings(dataPacking)

}
