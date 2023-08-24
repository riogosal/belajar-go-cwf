package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func HapusDataByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	path := "dbfilm.json"

	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
		return
	}

	var films []Film
	if err := json.Unmarshal(data, &films); err != nil {
		http.Error(w, "Parsing Data Gagal!", http.StatusInternalServerError)
		return
	}

	hapusID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID Tidak di temukan/Tidak sesuai", http.StatusBadRequest)
		return
	}

	indexDataFilm := -1
	for i, film := range films {
		if film.ID == hapusID {
			indexDataFilm = i
			break
		}
		fmt.Println(i)
	}

	if indexDataFilm == -1 {
		http.Error(w, "Film tidak ditemukan", http.StatusNotFound)
		return
	}

	films = append(films[:indexDataFilm], films[indexDataFilm+1:]...)

	newData, err := json.Marshal(films)
	if err != nil {
		http.Error(w, "Gagal menulis data", http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile(path, newData, 0644); err != nil {
		http.Error(w, "Gagal menulis data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Film berhasil dihapus")
}
