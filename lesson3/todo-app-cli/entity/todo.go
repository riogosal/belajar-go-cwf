package entity

import "fmt"

type Kegiatan struct {
	ID           int64
	Judul        string
	Deskripsi    string
	Selesai      bool
	KoleksiLabel []Label
}

func (kegiatan Kegiatan) PrintStuff() {
	fmt.Println(kegiatan.Judul)
	fmt.Println(kegiatan.Deskripsi)
	fmt.Println(kegiatan.KoleksiLabel)
}

type Label struct {
	Nama  string
	Warna string
}
