package main

import (
	"encoding/json"
	"fmt"
)

func display(db *VectorDB) {
	for _, entry := range db.Entries {
		displayEntry(entry)
	}
	fmt.Println()
}

func displayEntry(entry VectorEntry) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(entry.Data, &jsonData); err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
	jsonStr, _ := json.MarshalIndent(jsonData, "", "  ")
	fmt.Printf("Key: %s, Vector: %v, Data: %s\n", entry.Key, entry.Vector, string(jsonStr))
}

func JSONAsString(entry VectorEntry) string {
    var jsonData map[string]interface{}
    if err := json.Unmarshal(entry.Data, &jsonData); err != nil {
        fmt.Println("Error unmarshalling data:", err)
        return "" 
    }
    jsonStr, err := json.MarshalIndent(jsonData, "", "  ")
    if err != nil {
        fmt.Println("Error marshalling JSON:", err)
        return ""
    }
    return string(jsonStr)
}
