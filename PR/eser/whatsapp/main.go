package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"whatsapp-nautilus/domain"
	"whatsapp-nautilus/models"
	"whatsapp-nautilus/mongodb"
	"whatsapp-nautilus/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	r           domain.RawMaterialLotRepository
	rf          domain.RefinedmaterialRepository
	plc         domain.ProductLogContainersRepository
	p           domain.PackingRepository
	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)

// func init() {
// 	fmt.Println("Selamat Malam Chen Woo Fishery (Makassar) ini rangkmuman processing tanggal 2023-06-23")
// }

func Salam() string {
	now := time.Now()
	hour := now.Hour()

	if hour >= 6 && hour <= 9 {
		return "Selamat pagi"
	} else if hour >= 10 && hour <= 13 {
		return "Selamat siang"
	} else if hour >= 14 && hour <= 18 {
		return "Selamat Sore"
	} else {
		return "Selamat Malam"
	}
}

func main() {
	defer fmt.Println("\nSekedar informasinya untuk hari ini, Terima Kasih")

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
	p = repository.NewPacking(db)

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

	// inputTanggal := "2023-04-15"
	inputTanggal := "2023-06-26"

	t, err := time.Parse("2006-01-02", inputTanggal)
	if err != nil {
		fmt.Println("Parsing Tanggal gagal", err)
		return
	}

	startOfDayInt := t.UnixNano() / int64(time.Millisecond)

	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	endOfDayInt := endOfDay.UnixNano() / int64(time.Millisecond)

	// fmt.Println("Start of Day (millis):", startOfDayInt)
	// fmt.Println("End of Day (millis):", endOfDayInt)

	year := t.Year()
	month := t.Month()
	day := t.Day()

	fmt.Println(Salam(), "Chen Woo Fishery (Makassar), berikut rangkuman processing tanggal", day, month, year)

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

	resultPLC, err := plc.GetDataProductLogByDate(ctx, startOfDayInt, endOfDayInt)
	if err != nil {
		log.Fatal(err)
	}

	models.PrintProductLogContainerDataByDate(resultPLC)

	// ======================================================== //

	resultPacking, err := p.GetDataPackingByDate(ctx, startOfDayInt, endOfDayInt)
	if err != nil {
		log.Fatal(err)
	}

	models.PrintPackingDataByDate(resultPacking)
}
