package main

import (
	"app-api-movie/router"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	route.HandleFunc("/getMovies", router.JsonMovie)
	route.HandleFunc("/postMovies", router.JsonMovie)
	route.HandleFunc("/putMovies/{id}", router.JsonMovie)
	route.HandleFunc("/patchMovies/{id}", router.JsonMovie)

	server := http.Server{
		Addr:    ":8000",
		Handler: route,
	}
	fmt.Println("LOcalhost", server.Addr, "is ready")
	server.ListenAndServe()

}
