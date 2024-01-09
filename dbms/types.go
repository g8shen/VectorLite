package main

import (
	"encoding/json"
)

type Vector []float64

type VectorEntry struct {
	Vector Vector
	Key    string
	Data   json.RawMessage
}

type VectorDB struct {
	Entries []VectorEntry
}

type distanceEntry struct {
	entry    VectorEntry
	distance float64
}
