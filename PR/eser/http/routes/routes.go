package routes

import (
	"belajar-go-cwf/PR/eser/http/controller"
	"net/http"

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
		Addr:    ":8080",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
