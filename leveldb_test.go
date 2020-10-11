package main

import (
	"fmt"
	"testing"

	"github.com/golang/leveldb"
	"github.com/golang/leveldb/bloom"
	"github.com/golang/leveldb/db"
	"github.com/golang/leveldb/memfs"
)

func benchGolangLevelDB_Get(b *testing.B, options *db.Options) {
	d, err := leveldb.Open("testdata-leveldb", options)
	if err != nil {
		fmt.Printf("Open failed: %v", err)
	}
	defer d.Close()

	set := func(k, v []byte) error {
		return d.Set(k, v, nil)
	}
	get := func(k []byte) ([]byte, error) {
		return d.Get(k, nil)
	}

	benchmarkGet(b, set, get)
}

func BenchmarkGet_GolangLevelDB(b *testing.B) {
	b.Run("Memory", func(b *testing.B) {
		o := &db.Options{
			FileSystem:   memfs.New(),
			FilterPolicy: bloom.FilterPolicy(10),
		}
		benchGolangLevelDB_Get(b, o)
	})
	b.Run("SSD(Bloom:10)", func(b *testing.B) {
		o := &db.Options{
			FilterPolicy: bloom.FilterPolicy(10),
		}
		benchGolangLevelDB_Get(b, o)
	})
}
