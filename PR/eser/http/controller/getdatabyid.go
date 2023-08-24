package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func GetDataByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	path := "dbfilm.json"

	file, err := os.Open(path)
	if err != nil {
		http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var films []Film

	if err := json.NewDecoder(file).Decode(&films); err != nil && err != io.EOF {
		http.Error(w, "Parsing Data Gagal!", http.StatusInternalServerError)
		return
	}

	// Cari film berdasarkan ID
	var cariIDfilm *Film
	for _, film := range films {
		if strconv.Itoa(film.ID) == id {
			cariIDfilm = &film
			break
		}
	}

	if cariIDfilm == nil {
		http.Error(w, "Film tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cariIDfilm)
}
