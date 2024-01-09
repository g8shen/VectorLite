package main

func (db *VectorDB) Insert(vectorEntry VectorEntry) {
	db.Entries = append(db.Entries, vectorEntry)
}

func oneHotEncode(category string, categories []string) []float64 {
	vector := make([]float64, len(categories))
	for i, cat := range categories {
		if cat == category {
			vector[i] = 1
			break
		}
	}
	return vector
}
