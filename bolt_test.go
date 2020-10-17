package main

import (
	"log"
	"testing"
	//"github.com/dgraph-io/badger"
	//"github.com/boltdb/bolt"
	"github.com/boltdb/bolt"
)

func benchBolt_Get(b *testing.B, options *bolt.Options) {
	db, err := bolt.Open("testdata-bolt", 0600, options)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a writable transaction.
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	bkt_name := "MyBucket"
	bkt, err := tx.CreateBucketIfNotExists([]byte(bkt_name))
	if err != nil {
		log.Fatal(err)
	}

	bkt = tx.Bucket([]byte(bkt_name))
	set := bkt.Put

	get := func(key []byte) ([]byte, error) {
		var result []byte
		err := db.View(func(tx *bolt.Tx) error {
			bkt := tx.Bucket([]byte(bkt_name))
			result = bkt.Get(key)
			return nil
		})
		//fmt.Printf("Get(%s)==%s\n", key, result)
		return result, err
	}

	benchmarkGet(b, set, get, tx.Commit)
}

func BenchmarkGet_Bolt(b *testing.B) {
	b.Run("SSD", func(b *testing.B) {
		benchBolt_Get(b, nil)
	})
}
