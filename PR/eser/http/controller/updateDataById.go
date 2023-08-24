package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateDataByID(w http.ResponseWriter, r *http.Request) {
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

	updateID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	updateFilm := Film{}
	if err := json.NewDecoder(r.Body).Decode(&updateFilm); err != nil {
		http.Error(w, "Gagal membaca data", http.StatusBadRequest)
		return
	}

	boolUpdated := false
	for i, film := range films {
		if film.ID == updateID {
			films[i] = updateFilm
			boolUpdated = true
			break
		}
	}

	if !boolUpdated {
		http.Error(w, "Film tidak ditemukan", http.StatusNotFound)
		return
	}

	newData, err := json.Marshal(films)
	if err != nil {
		http.Error(w, "Parsing Data gagal", http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile(path, newData, 0644); err != nil {
		http.Error(w, "Gagal menyimpan data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Update Film Berhasil")
}
