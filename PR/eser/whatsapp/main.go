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
	rf          domain.RefinedmaterialRepository
	plc         domain.ProductLogContainersRepository
	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)

func init() {
	fmt.Println("Data Laporan Receiving Tgl : ", time.Now())
}

func main() {

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mongoclient, err = mongodb.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	/* Raw Material Lot */
	db := mongoclient.Database("nautilus")
	r = repository.NewRawMaterialLot(db)
	rf = repository.NewRefinedMaterial(db)
	plc = repository.NewProductLogContainers(db)

	c, _ := context.WithTimeout(ctx, 10*time.Second)
	defer mongoclient.Disconnect(c)
	err = mongoclient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// id := "64151aff954fea8407a5aa83"
	// mydata, err := r.GetData(ctx, id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(mydata.CetakStrRawmaterialLot())

	//============== Menampilkan data berdasarkan range Tanggal==================//

	inputTanggal := "2023-04-15"

	t, err := time.Parse("2006-01-02", inputTanggal)
	if err != nil {
		fmt.Println("Parsing Tanggal gagal", err)
		return
	}

	startOfDayInt := t.UnixNano() / int64(time.Millisecond)

	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	endOfDayInt := endOfDay.UnixNano() / int64(time.Millisecond)

	fmt.Println("Start of Day (millis):", startOfDayInt)
	fmt.Println("End of Day (millis):", endOfDayInt)

	resultData, err := r.GetDataByDate(ctx, startOfDayInt, endOfDayInt)
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range resultData {
		fmt.Println(data.CetakDataByDate())
	}

	//==============================================//
	resultRefined, err := rf.GetDataRefinedByDate(ctx, startOfDayInt, endOfDayInt)
	if err != nil {
		log.Fatal(err)
	}

	for _, dataRefined := range resultRefined {
		fmt.Println(dataRefined.PrintRefinedMaterialDataByDate())
	}

	//==================================================//

	resultProductLogContainer, err := plc.GetDataProductLogByDate(ctx, startOfDayInt, endOfDayInt)
	if err != nil {
		log.Fatal(err)
	}

	for _, dataProductLogContainer := range resultProductLogContainer {
		fmt.Println(dataProductLogContainer.PrintProductLogContainerDataByDate())
	}
}
