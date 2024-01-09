package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	crudDemo()
	retrevialDemo()
}

// Retrevial, kNN search, euclid + cosine demo
func retrevialDemo() {

	db1 := VectorDB{}
	err := db1.Load("../datastore/media.gob")
	if err != nil {
		fmt.Println("Load error:", err)
	}

	if len(db1.Entries) == 0 {
		fmt.Println("No data loaded from file, populating database with random dataset.")
		media := generateRandomMedia(1000)
		for _, media := range media {
			vectorEntry := createMediaVector(media)
			db1.Insert(vectorEntry)
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
	nearestNeighbours := db1.kNN(queryVector.Vector, k, euclideanDistance)

	fmt.Println("Shows/Movies most similar to the query:")
	for _, media := range nearestNeighbours {
		fmt.Printf("Key: %s, Vector: %v\n", media.Key, media.Vector)
	}

	nearestNeighbours = db1.kNN(queryVector.Vector, k, cosineSimilarity)

	fmt.Println("Shows/Movies most similar to the query:")
	for _, media := range nearestNeighbours {
		fmt.Printf("Key: %s, Vector: %v\n", media.Key, media.Vector)
	}

	err = db1.Save("../datastore/media.gob")
	if err != nil {
		fmt.Println("Save error:", err)
		return
	}
}

// Crud Demo
func crudDemo() {
	db2 := VectorDB{}
	db2.Clear()
	db2.Save("../datastore/dummy.gob")
	fmt.Println("Inserting Data:")
	for i := 0; i < 3; i++ {
		vector := Vector{float64(i), float64(i) * 2}
		data, _ := json.Marshal(map[string]string{"info": fmt.Sprintf("data %d", i)})
		entry := VectorEntry{Vector: vector, Key: fmt.Sprintf("key%d", i), Data: json.RawMessage(data)}
		db2.Insert(entry)
	}
	display(&db2)

	fmt.Println("\nUpdating 'key2':")
	newVector := Vector{20, 40}
	db2.Update("key2", newVector)
	display(&db2)

	fmt.Println("\nDeleting 'key2':")
	db2.Delete("key2")
	display(&db2)
	db2.Save("../datastore/dummy.gob")
}
