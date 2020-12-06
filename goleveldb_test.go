package main

import (
	"fmt"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func benchGoLevelDB_Get(b *testing.B, options *opt.Options) {
	b.StopTimer()
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
	benchmarkGet(b, set, get, nil)
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

func benchGoLevelDB_Range(b *testing.B, options *opt.Options) {
	b.StopTimer()
	db, err := leveldb.OpenFile("testdata-goleveldb", options)
	if err != nil {
		fmt.Printf("Open failed: %v", err)
	}
	defer db.Close()

	set := func(k, v []byte) error {
		return db.Put(k, v, nil)
	}
	newIterator := func(start, end []byte) iterator.Iterator {
		return db.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(end)}, nil)
	}

	benchmarkRange(b, set, nil, newIterator)
}

func BenchmarkRange_syndtrGoLevelDB(b *testing.B) {
	b.Run("SSD(Bloom:No)", func(b *testing.B) {
		benchGoLevelDB_Range(b, nil)
	})
}
