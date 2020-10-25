package main

import (
	"fmt"
	"testing"

	"github.com/prologic/bitcask"
)

func benchBitcask_Get(b *testing.B, options ...bitcask.Option) {
	db, err := bitcask.Open("testdata-bitcask", options...)
	if err != nil {
		fmt.Printf("Open failed: %v", err)
	}
	defer func() {
		db.DeleteAll()
		db.Close()
	}()

	set := func(k, v []byte) error {
		return db.Put(k, v)
	}
	get := func(k []byte) ([]byte, error) {
		return db.Get(k)
	}
	benchmarkGet(b, set, get, nil)
}

func BenchmarkGet_Bitcask(b *testing.B) {
	/*
		b.Run("SSD()", func(b *testing.B) {
			benchGoLevelDB_Get(b, nil)
		})
	*/

	b.Run("SSD()", func(b *testing.B) {
		o := []bitcask.Option{
			bitcask.WithMaxKeySize(32),
			bitcask.WithMaxValueSize(32),
		}

		benchBitcask_Get(b, o...)
	})
}
