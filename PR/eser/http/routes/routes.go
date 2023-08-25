package routes

import (
	"eser-http/controller"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func SetupRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/get", controller.ShowData)
	r.HandleFunc("/create-film", controller.CreateFilm)
	r.HandleFunc("/get-data/{id}", controller.GetDataByID)
	r.HandleFunc("/hapus-data/{id}", controller.HapusDataByID)
	r.HandleFunc("/update-data/{id}", controller.UpdateDataByID)

	http.Handle("/", r)

	server := http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: r,
	}

	fmt.Println("Running server")
	err := server.ListenAndServe()
	fmt.Println("Saya tidak akan di print")
	if err != nil {
		panic(err)
	}
}
