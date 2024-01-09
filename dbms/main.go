package main

import (
	"fmt"
)

func main() {

	//VECTOR RETREVIAL DEMO

	db := VectorDB{}

	err := db.Load("../datastore/media.gob")
	if err != nil {
		fmt.Println("Load error:", err)
	}

	if len(db.Entries) == 0 {
		fmt.Println("No data loaded from file, populating database with random dataset.")
		media := generateRandomMedia(1000)
		for _, media := range media {
			vectorEntry := createMediaVector(media)
			db.Insert(vectorEntry)
		}

	}

	queryMedia := Media{
		Title:           "Query Show", // Title is not used in the query, can be any string
		Type:            "Movie",      // Type of the show, either "Movie" or "Show"
		Rating:          4.0,          // Rating out of 5 (5 being the highest rating)
		Length:          0.6,          // Normalized length (e.g., 0 for short, 1 for long)
		Year:            0.8,          // Normalized release year (0 for older, 1 for recent)
		Popularity:      0.75,         // Normalized popularity (0 for least popular, 1 for most popular)
		CriticalAcclaim: 4.2,          // Critical acclaim out of 5 (5 being highly critically acclaimed)
		AudienceTarget:  0.5,          // Normalized target audience (e.g., 0 for kids, 1 for adults)
		Genre:           "Action",     // Genre of the show/movie
	}

	queryVector := createMediaVector(queryMedia)

	k := 5
	nearestNeighbours := db.kNN(queryVector.Vector, k, euclideanDistance)

	fmt.Println("Shows/Movies most similar to the query:")
	for _, media := range nearestNeighbours {
		fmt.Printf("Key: %s, Vector: %v\n", media.Key, media.Vector)
	}

	nearestNeighbours = db.kNN(queryVector.Vector, k, cosineSimilarity)

	fmt.Println("Shows/Movies most similar to the query:")
	for _, media := range nearestNeighbours {
		fmt.Printf("Key: %s, Vector: %v\n", media.Key, media.Vector)
	}

	err = db.Save("../datastore/media.gob")
	if err != nil {
		fmt.Println("Save error:", err)
		return
	}
	fmt.Printf("Number of entries in the database: %d\n", len(db.Entries))

}
