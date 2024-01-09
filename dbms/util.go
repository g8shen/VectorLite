package main

import (
	"encoding/json"
	"fmt"
)

func display(db *VectorDB) {
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
