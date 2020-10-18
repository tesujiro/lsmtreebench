package main

import (
	"fmt"
	"testing"
)

func benchGolangMap_Get(b *testing.B) {
	db := make(map[string][]byte, b.N)

	set := func(k, v []byte) error {
		db[string(k)] = v
		return nil
	}
	get := func(k []byte) ([]byte, error) {
		value, ok := db[string(k)]
		if !ok {
			fmt.Printf("map for %s not found", k)
			return value, fmt.Errorf("map for %s not found", k)
		}
		return value, nil

	}

	benchmarkGet(b, set, get, nil)
}

func BenchmarkGet_GolangMap(b *testing.B) {
	b.Run("Memory", func(b *testing.B) {
		benchGolangMap_Get(b)
	})
}
