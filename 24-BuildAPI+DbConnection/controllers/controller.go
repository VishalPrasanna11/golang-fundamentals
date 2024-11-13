package controllers

import (
	model "Build_API_DB_Connection/models"
	"Build_API_DB_Connection/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddMovieHandler handles the addition of a new movie to the database.
func AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Movie Handler")
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.AddMovie(movie)
	if err != nil {
		http.Error(w, "Failed to add movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

// GetAllMoviesHandler retrieves all movies from the database.
func GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Movies Handler")
	w.Header().Set("Content-Type", "application/json")

	movies, err := services.GetAllMovies()
	if err != nil {
		http.Error(w, "Failed to get movies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

// GetMovieByIDHandler retrieves a specific movie by its ID.
func GetMovieByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Use mux.Vars to get path parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Convert the id to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	movie, err := services.GetMovieById(objectID)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// UpdateMovieHandler updates an existing movie in the database.
func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.UpdateMovie(objectID, movie)
	if err != nil {
		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// DeleteMovieHandler deletes a movie from the database by its ID.
func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = services.DeleteMovie(objectID)
	if err != nil {
		http.Error(w, "Failed to delete movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func MarkMovieAsWatchedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = services.MarkMovieAsWatched(objectID)
	if err != nil {
		http.Error(w, "Failed to mark movie as watched", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
