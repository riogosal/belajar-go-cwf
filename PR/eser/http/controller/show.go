package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func ShowData(w http.ResponseWriter, r *http.Request) {
	path := "dbfilm.json"

	file, err := os.Open(path)
	if err != nil {
		http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var films *[]Film
	if err := json.NewDecoder(file).Decode(&films); err != nil && err != io.EOF {
		http.Error(w, "Parsing Data Gagal!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(films); err != nil {
		http.Error(w, "Gagal mengirimkan data", http.StatusInternalServerError)
		return
	}
}
