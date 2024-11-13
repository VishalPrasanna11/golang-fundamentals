package routes

import (
	"Build_API_DB_Connection/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	content := "Health Check"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, content) // Write the content to the response
	}).Methods("GET")
	router.HandleFunc("/movies", controllers.GetAllMoviesHandler).Methods("GET")
	router.HandleFunc("/movies", controllers.AddMovieHandler).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.GetMovieByIDHandler).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovieHandler).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovieHandler).Methods("DELETE")
	router.HandleFunc("/watch/{id}", controllers.MarkMovieAsWatchedHandler).Methods("PUT")

	return router
}
