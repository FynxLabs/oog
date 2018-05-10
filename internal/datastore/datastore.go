package datastore

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

// Load Brain
func Load() {
	// Open the bolt.db data file in your brain directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("./brain/bolt.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

// Save - Save stuff to datastore/brain
func Save(store string, key string, value string) {
}

// Delete - Delete stuff from datastore/brain
func Delete(store string, key string) {
}

// Query - Find stuff in datastore/brain
func Query(values ...string) {
	if len(values) == 2 {

	} else {

	}
}
