package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Film struct {
	ID        int
	Judul     string
	Deskripsi string
	Studio    string
	Rating    float64
	Durasi    int
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {

	path := "dbfilm.json"

	/**
	os.O_RDWR (membaca dan menulis)
	os.O_CREATE (buat file jika tidak ada)
	os.ModePerm (izin akses read, write, dan execute kepada pemilik file)
	*/

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var films []Film

	if err := json.NewDecoder(file).Decode(&films); err != nil && err != io.EOF { // kalau di tengah file terjadi error saat baca file maka responnya Gagal! end of file
		http.Error(w, "Parsing Data Gagal!", http.StatusInternalServerError)
		return
	}

	var newFilm Film
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Gagal membaca data", http.StatusBadRequest)
		return
	}
	// fmt.Println(string(body))
	if err := json.Unmarshal(body, &newFilm); err != nil {
		http.Error(w, "Parsing Data Gagal!", http.StatusBadRequest)
		return
	}

	for _, film := range films {
		if film.ID == newFilm.ID {
			http.Error(w, "ID sudah ada dalam database", http.StatusBadRequest)
			return
		}
	}

	films = append(films, newFilm)

	//Penting !! https://stackoverflow.com/questions/44416645/how-do-i-truncate-and-completely-rewrite-a-file-without-having-leading-zeros

	file.Seek(0, 0)
	file.Truncate(0)

	if err := json.NewEncoder(file).Encode(films); err != nil {
		http.Error(w, "Gagal menulis data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("\nID       : %d\n", newFilm.ID)
	fmt.Printf("Judul    : %s\n", newFilm.Judul)
	fmt.Printf("Deskripsi: %s\n", newFilm.Deskripsi)
	fmt.Printf("Studio   : %s\n", newFilm.Studio)
	fmt.Printf("Rating   : %.2f\n", newFilm.Rating)
	fmt.Printf("Durasi   : %d\n", newFilm.Durasi)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Film berhasil dibuat")
}
