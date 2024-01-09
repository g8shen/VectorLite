package main

import (
	"math"
	"sort"
)

func (db *VectorDB) SearchByKey(key string) *VectorEntry {
	for _, entry := range db.Entries {
		if entry.Key == key {
			return &entry
		}
	}
	return nil
}

func euclideanDistance(a, b Vector) float64 {
	var sum float64
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func cosineSimilarity(a, b Vector) float64 {
	var dotProduct, normA, normB float64
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

func (db *VectorDB) kNN(query Vector, k int, distanceMetric func(Vector, Vector) float64) []VectorEntry {
	var distances []distanceEntry
	for _, entry := range db.Entries {
		dist := distanceMetric(query, entry.Vector)
		distances = append(distances, distanceEntry{entry, dist})
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	var results []VectorEntry
	for i := 0; i < k && i < len(distances); i++ {
		results = append(results, distances[i].entry)
	}

	return results
}


