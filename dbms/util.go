package main

import (
	"encoding/json"
	"fmt"
)

//	func printDatabaseEntries(db *VectorDB) {
//		fmt.Printf("Number of entries in the database: %d\n", len(db.Entries))
//		for _, entry := range db.Entries {
//			var media Media
//			if err := json.Unmarshal(entry.Data, &media); err != nil {
//				fmt.Println("Error unmarshalling data:", err)
//				continue
//			}
//			fmt.Printf("Media: %+v\n", media)
//		}
//	}
func printDatabaseEntries(db *VectorDB) {
	for _, entry := range db.Entries {
		var jsonData map[string]interface{}
		if err := json.Unmarshal(entry.Data, &jsonData); err != nil {
			fmt.Println("Error unmarshalling data:", err)
			continue
		}
		jsonStr, _ := json.MarshalIndent(jsonData, "", "  ")
		fmt.Printf("Key: %s, Vector: %v, Data: %s\n", entry.Key, entry.Vector, string(jsonStr))
	}
	fmt.Println()
}
