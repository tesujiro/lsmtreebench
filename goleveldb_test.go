package main

import (
	"fmt"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func benchGoLevelDB_Get(b *testing.B, options *opt.Options) {
	db, err := leveldb.OpenFile("testdata-goleveldb", options)
	if err != nil {
		fmt.Printf("Open failed: %v", err)
	}
	defer db.Close()

	set := func(k, v []byte) error {
		return db.Put(k, v, nil)
	}
	get := func(k []byte) ([]byte, error) {
		return db.Get(k, nil)
	}
	benchmarkGet(b, set, get)
}

func BenchmarkGet_syndtrGoLevelDB(b *testing.B) {
	b.Run("SSD(Bloom:No)", func(b *testing.B) {
		benchGoLevelDB_Get(b, nil)
	})

	b.Run("SSD(Bloom:4)", func(b *testing.B) {
		o := &opt.Options{
			Filter: filter.NewBloomFilter(4),
		}
		benchGoLevelDB_Get(b, o)
	})
	b.Run("SSD(Bloom:10)", func(b *testing.B) {
		o := &opt.Options{
			Filter: filter.NewBloomFilter(10),
		}
		benchGoLevelDB_Get(b, o)
	})
}
