package main

import (
	"fmt"
	"testing"

	"github.com/golang/leveldb"
	"github.com/golang/leveldb/db"
	"github.com/golang/leveldb/memfs"
)

func benchGolangLevelDB_Get(b *testing.B, options *db.Options) {
	b.StopTimer()

	b.Run("SSD", func(b *testing.B) {
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
	})
}

func BenchmarkGet_GolangLevelDB(b *testing.B) {
	b.Run("Memory", func(b *testing.B) {
		o := &db.Options{
			FileSystem: memfs.New(),
		}
		benchGolangLevelDB_Get(b, o)
	})
	b.Run("SSD", func(b *testing.B) {
		o := &db.Options{}
		benchGolangLevelDB_Get(b, o)
	})
}
