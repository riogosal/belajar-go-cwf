package controller

import (
	"net/http"
)

// Definisikan tipe Film dan slice Films

// type Films []Film

// Definisikan fungsi findFilmByID
// func findFilmByID(id int) (*Film, error) {
// Baca data dari file dbfilm.json
// data, err := os.ReadFile("dbfilm.json")
// if err != nil {
// 	return nil, err
// }

// // Parse data menjadi slice Films
// var films Films
// if err := json.Unmarshal(data, &films); err != nil {
// 	return nil, err
// }

// // Cari film berdasarkan ID
// for _, film := range films {
// 	if film.ID == id {
// 		return &film, nil
// 	}
// }

// // Jika ID tidak ditemukan, kembalikan error
// return nil, errors.New("Film tidak ditemukan")
// }

func UpdateFilm(w http.ResponseWriter, r *http.Request) {

}

// filmToUpdate, err := findFilmByID(Newid)
// if err != nil {
// 	http.Error(w, "Film tidak ditemukan", http.StatusNotFound)
// 	return
// }

// // Membaca data baru dari body permintaan
// var updatedFilm Film
// err = json.NewDecoder(r.Body).Decode(&updatedFilm)
// if err != nil {
// 	http.Error(w, "Gagal membaca data", http.StatusBadRequest)
// 	return
// }

// // Memperbarui data film
// filmToUpdate.Judul = updatedFilm.Judul
// filmToUpdate.Deskripsi = updatedFilm.Deskripsi
// filmToUpdate.Studio = updatedFilm.Studio
// filmToUpdate.Rating = updatedFilm.Rating
// filmToUpdate.Durasi = updatedFilm.Durasi

// // Simpan perubahan ke database Anda (misalnya, ke file JSON)

// // Respon sukses
// w.WriteHeader(http.StatusOK)
// fmt.Fprintln(w, "Film berhasil diupdate")
// }
