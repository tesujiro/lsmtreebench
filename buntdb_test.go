package main

import (
	"log"
	"testing"

	"github.com/tidwall/buntdb"
)

func benchBuntDB_Get(b *testing.B, path string) {
	b.StopTimer()
	db, err := buntdb.Open(path)
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

	set := func(key, value []byte) error {
		_, _, err := tx.Set(string(key), string(value), nil)
		return err
	}

	get := func(key []byte) ([]byte, error) {
		var value string
		err := db.View(func(tx *buntdb.Tx) error {
			value, err = tx.Get(string(key))
			return err
		})
		//fmt.Printf("key=%s\tvalue=%s\n", key, value)
		return []byte(value), err
	}

	benchmarkGet(b, set, get, tx.Commit)
}

func BenchmarkGet_BuntDB(b *testing.B) {
	b.Run("Memory", func(b *testing.B) {
		benchBuntDB_Get(b, ":memory:")
	})
	b.Run("SSD", func(b *testing.B) {
		benchBuntDB_Get(b, "testdata-buntdb")
	})
}
