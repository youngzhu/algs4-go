package client

import (
	"fmt"

	"github.com/youngzhu/algs4-go/searching"
	"github.com/youngzhu/algs4-go/util"
)

// FrequencyCounter is a symbol-table client that finds the number of occurrences
// of each string (having at least as many characters as a given threshold length)
// in a squence of strings, then iterates through the keys to find the one that
// occurs the most frequently.

func FrequencyCounter(path string, minLen int) *FrequencyResult {
	st := searching.NewSequentialSearchST()

	// compute frequency counts
	distinct, words := 0, 0
	in := util.NewInReadWords(path)
	for !in.IsEmpty() {
		word := in.ReadString()
		if len(word) < minLen {
			continue
		}
		words++
		key := searching.StringKey(word)
		if st.Contains(key) {
			value := st.Get(key).(int)+1
			st.Put(key, value)
		} else {
			st.Put(key, 1)
			distinct++
		}
	}

	// find a word with the highest frequency count
	var max searching.STKey = searching.StringKey("")
	st.Put(max, 0)
	for _, w := range st.Keys() {
		wInt := st.Get(w).(int)
		maxInt := st.Get(max).(int)
		if wInt > maxInt {
			max = w
		}
	}

	return &FrequencyResult{max, st.Get(max), minLen, words, distinct}
}

type FrequencyResult struct {
	max searching.STKey
	frequency searching.STValue
	minLen, words, distinct int
}

func (r FrequencyResult) String() string {
	const str = "high-frequency word: %v, frequency: %v, minLen: %d, total words: %d, distinct words: %d"
	return fmt.Sprintf(str, r.max, r.frequency, r.minLen, r.words, r.distinct)
}