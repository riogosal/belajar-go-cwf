package router

import (
	"app-api-movie/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func JsonMovie(w http.ResponseWriter, r *http.Request) {
	var film model.Movie
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(model.Film)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	case "POST":

		err := json.NewDecoder(r.Body).Decode(&film)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		model.Film = append(model.Film, film)
		encoder := json.NewEncoder(w)
		encoder.Encode(model.Film)

	case "PUT":
		vars := mux.Vars(r)
		key := vars["id"]
		fmt.Fprintf(w, "Key: "+key)
		k, err := strconv.Atoi(key)
		fmt.Println(k)
		if err != nil {
			panic("error")
		}

		// Loop over all of our Articles
		// if the article.Id equals the key we pass in
		// return the article encoded as JSON
		for i, v := range model.Film {
			if v.Id == k {
				filmId := model.Film[i]
				json.NewEncoder(w).Encode(filmId)
			}
		}

	case "PATCH":
		vars := mux.Vars(r)
		key := vars["id"]
		k, err := strconv.Atoi(key)
		fmt.Println(k)
		if err != nil {
			panic("error")
		}

		// Loop over all of our Articles
		// if the article.Id equals the key we pass in
		// return the article encoded as JSON
		for i, v := range model.Film {
			if v.Id == k {
				filmId := model.Film[i]

				err := json.NewDecoder(r.Body).Decode(&filmId)
				if err != nil {
					http.Error(w, "Bad request", http.StatusBadRequest)
					return
				}
				json.NewEncoder(w).Encode(model.Film)
			}
		}

	}

}
