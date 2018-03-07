package oog

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"net/http"
)

// fun brain()
func brain() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("oog_brain.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	db, err := bolt.Open("oog_brain.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
