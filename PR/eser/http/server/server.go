package server

import (
	"fmt"
	"net/http"
)

var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello There, Eser's Here!")
}

func Server() {
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	fmt.Printf("Server berjalan di http://localhost%v\n", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
