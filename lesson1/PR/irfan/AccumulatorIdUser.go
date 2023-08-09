package main

import (
	"fmt"
	"strconv"
	"time"
)

func heartBeat() {
	time.Sleep(time.Second * 3)
	fmt.Println("Program telah dipersiapkan...")

}

func inputDataIdUsers() {
	time.Sleep(time.Second * 5)
	var nilai []int

	type id string
	type start bool
	var mulai start = true
	for mulai {
		fmt.Println("Masukkan User Id anda berupa angka, tekan (e) exit : ")
		var UserId id
		fmt.Scanln(&UserId)

		if UserId == "e" {
			defer fmt.Println("program Selesai... Arigatou !!")
			break

		} else if data, err := strconv.Atoi(string(UserId)); err == nil {
			nilai = append(nilai, data)
			fmt.Println("nilai anda sudah dimasukan")
		} else {
			fmt.Println("nilai yang anda masukkan salah")
		}
	}
	fmt.Println("data Id yang Anda input adalah ", nilai)
	ResultData(nilai)
	dataMedian(nilai)

}

func dataMedian(sliceData []int) {
	if len(sliceData)%2 == 0 {
		medianGenapSatu := (len(sliceData) / 2)
		medianGenapSatu = sliceData[medianGenapSatu-1]
		medianGenapDua := (len(sliceData) / 2) + 1
		medianGenapDua = sliceData[medianGenapDua-1]

		ResultMedian := float32(medianGenapSatu) / float32(medianGenapDua)

		fmt.Printf("Hasil dari median User Id adalah = %.2f\n", ResultMedian)
	} else {
		medianGanjil := (len(sliceData) + 1) / 2
		ResultMedian := medianGanjil
		fmt.Printf("Hasil dari median User Id adalah = ", ResultMedian)
	}

}

func ResultData(sliceData []int) {
	var sum int = 0
	for _, v := range sliceData {
		sum += v
	}
	mean := float32(sum) / float32(len(sliceData))
	fmt.Println("hasil dari penjumlahan nilai Id = ", sum)
	fmt.Printf("hasil dari nilai rata-rata Id = %.2f\n", float32(mean))
}

func main() {
	fmt.Println("Go runs Program akumulasi Id Users On ")
	heartBeat()
	inputDataIdUsers()

}
