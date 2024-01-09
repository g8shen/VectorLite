package main

import (
	"encoding/json"
	"fmt"
)

func printDatabaseEntries(db *VectorDB) {
	fmt.Printf("Number of entries in the database: %d\n", len(db.Entries))
	for _, entry := range db.Entries {
		var media Media
		if err := json.Unmarshal(entry.Data, &media); err != nil {
			fmt.Println("Error unmarshalling data:", err)
			continue
		}
		fmt.Printf("Media: %+v\n", media)
	}
}
