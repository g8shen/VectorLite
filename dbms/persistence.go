package main

import (
	"encoding/gob"
	"os"
)

func (db *VectorDB) Save(filename string) error {
	file, err := os.Create(filename) // os.Create overwrites any existing file
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(db.Entries) // Encode only the current db.Entries
	if err != nil {
		return err
	}

	return nil
}

func (db *VectorDB) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&db.Entries)
	if err != nil {
		return err
	}

	return nil
}

func (db *VectorDB) Clear() {
	db.Entries = []VectorEntry{}
}
