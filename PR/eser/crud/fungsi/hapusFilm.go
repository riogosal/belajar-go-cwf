package fungsi

import (
	"belajar-go-cwf/PR/eser/crud/model"
	"fmt"
)

func HapusDataFilm(koleksi_film *[]model.Film) {
	fmt.Print("Masukkan ID Film yang Ingin Dihapus: ")
	var idHapus int
	_, err := fmt.Scan(&idHapus)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mencari film berdasarkan ID
	index := -1
	for i, film := range *koleksi_film {
		if film.ID == idHapus {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Film dengan ID", idHapus, "tidak ditemukan.")
		return
	}
	*koleksi_film = append((*koleksi_film)[:index], (*koleksi_film)[index+1:]...)
	fmt.Println("Film dengan ID", idHapus, "telah dihapus.")
}
