package netflixRouter

import (
	netflixController "github.com/Ashmit-05/mongodbapi/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",netflixController.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",netflixController.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",netflixController.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",netflixController.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/delete-all-movies",netflixController.DeleteAllMovies).Methods("DELETE")

	return router
}