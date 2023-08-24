package main

import (
	"app-api-movie/model"
	"app-api-movie/router"
	"fmt"
	"net/http"
)

func main() {
	var film []model.Movie
	fmt.Println(film)
	route := http.NewServeMux()
	route.HandleFunc("/getMovies", router.JsonMovie)
	route.HandleFunc("/postMovies", router.JsonMovie)

	server := http.Server{
		Addr:    ":8000",
		Handler: route,
	}
	fmt.Println("Localhost", server.Addr, "is ready")

	// Load data to model.Movies
	server.ListenAndServe()

}
