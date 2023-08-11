package main

import (
	"todo-app-cli/entity"
)

func main() {
	kegiatan_nomor_1 := entity.Kegiatan{
		ID:        1,
		Judul:     "Mandi",
		Deskripsi: "Mandi adalah aktivitas setiap pagi",
		Selesai:   false,
		KoleksiLabel: []entity.Label{
			{
				Nama:  "Harian",
				Warna: "biru",
			},
			{
				Nama:  "PENTING",
				Warna: "merah",
			},
		},
	}

	// fmt.Println(kegiatan_nomor_1)
	// Akses property kegiatan nomor 1
	// fmt.Println("Kegiatan pertama adalah: " + kegiatan_nomor_1.Judul)
	// fmt.Println("Deskripsi kegiatan nomor 1: " + kegiatan_nomor_1.Deskripsi)
	// fmt.Println("Label di kegiatan nomor 1: ", kegiatan_nomor_1.KoleksiLabel)

	// Kamu bisa buat kegiatan nomor 2
	kegiatan_nomor_2 := entity.Kegiatan{
		ID:           2,
		Judul:        "Pergi Kerja",
		Deskripsi:    "Aktivitas cari uang",
		Selesai:      false,
		KoleksiLabel: []entity.Label{},
	}

	// Akses property kegiatan nomor 2
	// fmt.Println("Kegiatan kedua adalah: " + kegiatan_nomor_2.Judul)
	// fmt.Println("Deskripsi kegiatan nomor 2: " + kegiatan_nomor_2.Deskripsi)
	// fmt.Println("Label di kegiatan nomor 2: ", kegiatan_nomor_2.KoleksiLabel)

	// Put inside collection and loop
	/* kode PHP
	foreach($kegiatanHariHari as $index => $kegiatan) {
		var_dump("Loop nomor", $index, $kegiatan)
		var_dump($kegiatan)
	}
	*/
	kegiatan_hari_hari := []entity.Kegiatan{kegiatan_nomor_1, kegiatan_nomor_2}
	for _, kegiatan := range kegiatan_hari_hari {
		// fmt.Println("Kegiatan" + fmt.Sprint(index+1) + "adalah: " + kegiatan.Judul)
		// fmt.Println(fmt.Sprintf("Deskripsi kegiatan nomor %d: ", index+1) + kegiatan.Deskripsi)
		// fmt.Println(fmt.Sprintf("Label kegiatan nomor %d: ", index+1), kegiatan.KoleksiLabel)
		kegiatan.PrintStuff()
	}
}
