package services

import (
	"Build_API_DB_Connection/config"
	model "Build_API_DB_Connection/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMovies() ([]model.Netflix, error) {
	// Use the global MongoDB client from config
	client := config.GetMongoClient()

	// Get the collection from the database
	collection := client.Database("GoLang").Collection("Netflix")

	// Define a slice to store the movies
	var movies []model.Netflix

	// Query the database to get all movies
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not fetch movies: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Iterate over the cursor and decode the documents into the movies slice
	for cursor.Next(context.TODO()) {
		var movie model.Netflix
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	// Check if there were any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %v", err)
	}

	return movies, nil
}

func AddMovie(movie model.Netflix) error {
	// Use the global MongoDB client from config
	client := config.GetMongoClient()

	// Get the collection from the database
	collection := client.Database("GoLang").Collection("Netflix")

	// Insert the movie document into the collection
	_, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return fmt.Errorf("could not insert movie: %v", err)
	}

	return nil
}

func UpdateMovie(id primitive.ObjectID, movie model.Netflix) error {
	// Use the global MongoDB client from config
	client := config.GetMongoClient()

	// Get the collection from the database
	collection := client.Database("GoLang").Collection("Netflix")

	// Update the movie document in the collection
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": movie})
	if err != nil {
		return fmt.Errorf("could not update movie: %v", err)
	}

	return nil
}

func DeleteMovie(id primitive.ObjectID) error {
	// Use the global MongoDB client from config
	client := config.GetMongoClient()

	// Get the collection from the database
	collection := client.Database("GoLang").Collection("Netflix")

	// Delete the movie document from the collection
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("could not delete movie: %v", err)
	}

	return nil
}

func GetMovieById(id primitive.ObjectID) (model.Netflix, error) {
	// Use the global MongoDB client from config
	client := config.GetMongoClient()

	// Get the collection from the database
	collection := client.Database("GoLang").Collection("Netflix")

	// Define a movie variable to store the result
	var movie model.Netflix

	// Query the database to get the movie by ID
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		return model.Netflix{}, fmt.Errorf("could not find movie: %v", err)
	}
	fmt.Println(`This is from Get Movie By Id`, movie)
	return movie, nil
}

func MarkMovieAsWatched(id primitive.ObjectID) error {
	// Use the global MongoDB client from config
	client := config.GetMongoClient()

	// Get the collection from the database
	collection := client.Database("GoLang").Collection("Netflix")

	// Update the movie document in the collection to mark it as watched
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"watched": true}})
	if err != nil {
		return fmt.Errorf("could not mark movie as watched: %v", err)
	}

	return nil
}
