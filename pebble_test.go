package main

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
)

func benchCockloacdbPebble_Get(b *testing.B, options *pebble.Options) {
	db, err := pebble.Open("testdata-pebble", options)
	if err != nil {
		fmt.Printf("Open failed: %v", err)
	}
	defer db.Close()

	set := func(k, v []byte) error {
		//return db.Set(k, v, pebble.Sync)
		return db.Set(k, v, pebble.NoSync)
	}
	get := func(k []byte) ([]byte, error) {
		value, _, err := db.Get(k)
		//fmt.Printf("Get key=%s\tvalue=%s\n", k, value)
		return value, err
	}

	benchmarkGet(b, set, get, nil)
}

func BenchmarkGet_pebble(b *testing.B) {
	b.Run("Memory", func(b *testing.B) {
		o := &pebble.Options{
			FS: vfs.NewMem(),
		}
		benchCockloacdbPebble_Get(b, o)
	})
	b.Run("SSD", func(b *testing.B) {
		o := &pebble.Options{}
		benchCockloacdbPebble_Get(b, o)
	})
	/*
		b.Run("SSD(Bloom:10)", func(b *testing.B) {
			o := &pebble.Options{}
			benchCockloacdbPebble_Get(b, o)
		})
	*/
}
