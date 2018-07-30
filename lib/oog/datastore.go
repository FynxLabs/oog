package oog

// import (
// 	"fmt"

// 	bolt "github.com/coreos/bbolt"
// )

// // Save - Save stuff to datastore/brain
// func dataSave(*brain, bucket string, store string, key string, value string) {
// 	brain.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucket([]byte(bucket))
// 		if err != nil {
// 			return fmt.Errorf("create bucket: %s", err)
// 		}
// 		return nil
// 	})
// }

// func createBucket(*brain, bucket string) {
// 	brain.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucket([]byte(bucket))
// 		if err != nil {
// 			return fmt.Errorf("create bucket: %s", err)
// 		}
// 		return nil
// 	})
// }

// // Delete - Delete stuff from datastore/brain
// func dataDelete(store string, key string) {
// }

// // Query - Find stuff in datastore/brain
// func dataQuery(values ...string) {
// 	if len(values) == 2 {

// 	} else {

// 	}
// }
