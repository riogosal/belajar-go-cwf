package main

import "fmt"

type DataKelas struct {
	Kelas string
	Nilai int
}

func main() {
	dataCounts := make(map[string]int)
	data := []DataKelas{{
		Kelas: "sains",
		Nilai: 90,
	},
		{
			Kelas: "math",
			Nilai: 100,
		},
		{
			Kelas: "Fisika",
			Nilai: 60,
		},
		{
			Kelas: "Fisika",
			Nilai: 30,
		},
		{
			Kelas: "sains",
			Nilai: 80,
		},
		{
			Kelas: "math",
			Nilai: 20,
		},
		{
			Kelas: "sejarah",
			Nilai: 65,
		},
	}

	for _, v := range data {

		// fmt.Print(dataCounts[v.Kelas])
		// dataCounts[v.Kelas] = 1
		if dataCounts[v.Kelas] == 0 {
			dataCounts[v.Kelas] = 1
			fmt.Println(dataCounts[v.Kelas])

		} else {
			dataCounts[v.Kelas]++
			fmt.Println(dataCounts[v.Kelas])
		}
	}
	fmt.Println(dataCounts)
}
