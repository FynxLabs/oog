package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

// fun brain()
func brain() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("oog_brain.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
