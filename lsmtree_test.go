package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/syndtr/goleveldb/leveldb/iterator"
)

func benchmarkGet(b *testing.B, set func([]byte, []byte) error, get func([]byte) ([]byte, error), commit func() error) {
	//b.ResetTimer()
	b.StopTimer()
	//fmt.Printf("b.N:%v\n", b.N)
	size := b.N
	keys := make([]string, size)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%.8d", i)
		val := fmt.Sprintf("val%.8d", i)
		keys[i] = key

		err := set([]byte(key), []byte(val))
		if err != nil {
			fmt.Printf("Set(%q): %v\n", key, err)
			return
		}
	}
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	if commit != nil {
		err := commit()
		if err != nil {
			fmt.Printf("Commit: %v\n", err)
			return
		}
	}

	//b.ResetTimer()
	b.StartTimer()

	for _, key := range keys {
		g, err := get([]byte(key))
		if err != nil {
			fmt.Printf("Get(%q): %v\n", key, err)
			return
		}
		//fmt.Printf("Get(%q): %s\n", key, g)
		_ = g
	}
	b.StopTimer()
}

func benchmarkRange(b *testing.B, set func([]byte, []byte) error, commit func() error, newIterator func(start, end []byte) iterator.Iterator) {
	//b.ResetTimer()
	b.StopTimer()
	//fmt.Printf("b.N:%v\n", b.N)
	size := b.N
	keys := make([]string, size)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%.8d", i)
		val := fmt.Sprintf("val%.8d", i)
		keys[i] = key

		err := set([]byte(key), []byte(val))
		if err != nil {
			fmt.Printf("Set(%q): %v\n", key, err)
			return
		}
	}
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	if commit != nil {
		err := commit()
		if err != nil {
			fmt.Printf("Commit: %v\n", err)
			return
		}
	}

	//b.ResetTimer()
	b.StartTimer()

	key_start := []byte(fmt.Sprintf("key%.8d", 0))
	key_end := []byte(fmt.Sprintf("key%.8d", b.N))
	iter := newIterator(key_start, key_end)
	for iter.Next() {
		// Use key/value.
		k := iter.Key()
		v := iter.Value()
		_, _ = k, v
		//fmt.Printf("key=%s\tvalue=%s\n", k, v)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Printf("Release: %v\n", err)
		return
	}

	b.StopTimer()
}
