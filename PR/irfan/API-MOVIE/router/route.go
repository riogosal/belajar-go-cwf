package router

import (
	"app-api-movie/model"
	"encoding/json"
	"io"
	"net/http"
)

func JsonMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(model.Movies)
		return
		// fmt.Fprintf(w, "response: %v", model.Movies)
	case "POST":
		var film model.Movie
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &film)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// fmt.Println(body)
		encoder := json.NewEncoder(w)
		encoder.Encode(film)

		model.Movies = append(model.Movies, film)
		json.NewEncoder(w).Encode(map[string]any{
			"status": "OK",
		})
		return

	}

}
