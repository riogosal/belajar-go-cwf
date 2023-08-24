package router

import (
	"app-api-movie/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func JsonMovie(w http.ResponseWriter, r *http.Request) {
	var films model.Movie
	tampungData := map[string]model.Movie{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "response: %v", tampungData)
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &films)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		fmt.Println(body)
		encoder := json.NewEncoder(w)
		encoder.Encode(films)

		tampungData[films.Title] = films
		fmt.Fprintf(w, "responses 1: %v", films)
		fmt.Fprintf(w, "responses 2: %v", tampungData)

	}

}
