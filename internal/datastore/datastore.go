package datastore

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

// Default brain session
func Load() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("oog_brain.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

// Save - Save stuff to datastore/brain
func Save() {
}

// Delete - Delete stuff from datastore/brain
func Delete() {
}

// Query - Find stuff in datastore/brain
func Query() {
}
