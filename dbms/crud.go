package main

func (db *VectorDB) Insert(vectorEntry VectorEntry) {
	db.Entries = append(db.Entries, vectorEntry)
}

func (db *VectorDB) Delete(key string) {
	var newEntries []VectorEntry
	for _, entry := range db.Entries {
		if entry.Key != key {
			newEntries = append(newEntries, entry)
		}
	}
	db.Entries = newEntries
}

func (db *VectorDB) Update(key string, newVector Vector) {
	for i, entry := range db.Entries {
		if entry.Key == key {
			db.Entries[i].Vector = newVector
			return
		}
	}
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
