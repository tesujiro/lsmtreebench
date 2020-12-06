package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/dgraph-io/badger"
)

func benchBadgerDB_Get(b *testing.B, options badger.Options) {
	b.StopTimer()
	db, err := badger.Open(options)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a writable transaction.
	txn := db.NewTransaction(true)
	defer txn.Discard()

	var set func([]byte, []byte) error
	set = func(key, value []byte) error {
		err := txn.Set(key, value)
		if err == badger.ErrTxnTooBig {
			fmt.Println("badger.ErrTxnTooBig")
			txn = db.NewTransaction(true)
			return set(key, value)
		}
		return err
	}

	get := func(key []byte) ([]byte, error) {
		var result []byte
		err := db.View(func(txn *badger.Txn) error {
			item, err := txn.Get(key)
			if err != nil {
				return err
			}
			return item.Value(func(val []byte) error {
				copy(result, val)
				return nil
			})
		})
		return result, err
	}

	benchmarkGet(b, set, get, txn.Commit)
}

func BenchmarkGet_BadgerDB(b *testing.B) {
	/*
		b.Run("Memory", func(b *testing.B) {
			o := badger.DefaultOptions("").WithInMemory(true).WithMaxTableSize(64 << 22).WithLoggingLevel(badger.WARNING)
			//o.Logger = nil1
			benchBadgerDB_Get(b, o)
		})
	*/
	b.Run("SSD", func(b *testing.B) {
		//o := badger.DefaultOptions("testdata-badger").WithMaxTableSize(64 << 22).WithLoggingLevel(badger.WARNING)
		o := badger.DefaultOptions("testdata-badger").WithMaxTableSize(64 << 24).WithLoggingLevel(badger.WARNING)
		benchBadgerDB_Get(b, o)
	})
}
