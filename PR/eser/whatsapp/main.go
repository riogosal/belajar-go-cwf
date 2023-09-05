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
	rawMaterial := mongoclient.Database("nautilus").Collection("raw_material_lots")
	r = repository.NewRawMaterialLot(rawMaterial, ctx)

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

	// utcTime := time.Now().UTC()
	mksTime, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		fmt.Println("Gagal memuat zona waktu:", err)
		return
	}

	/*
		localTime := utcTime.In(mksTime)
		startOfDay := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), 0, 0, 0, 0, mksTime)
		endOfDay := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), 23, 59, 59, 999999999, mksTime)
	*/

	// Input tanggal manual, karena dalam database tidak ada tanggal hari ini.
	startOfDay := time.Date(2023, time.March, 23, 0, 0, 0, 0, mksTime)
	endOfDay := time.Date(2023, time.March, 23, 23, 59, 59, 999999999, mksTime)

	startOfDayInt := startOfDay.UnixMilli()
	endOfDayInt := endOfDay.UnixMilli()

	resultData, err := r.GetDataByDate(ctx, startOfDayInt, endOfDayInt)
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range resultData {
		fmt.Println(data.CetakDataByDate())
	}
}
