package routes

import (
	"fmt"
	"go-movies-crud/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r *mux.Router

func InitializeRouter() {
	r = mux.NewRouter()
	r.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controller.AddMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")
}

func StartServer() {
	fmt.Printf("Server started on port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
