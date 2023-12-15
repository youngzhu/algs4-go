package searching_test

import (
	"fmt"
	// "log"

	"github.com/youngzhu/algs4-go/searching"
	"github.com/youngzhu/algs4-go/testutil"
)

// FrequencyCounter is a symbol-table client that finds the number of occurrences
// of each string (having at least as many characters as a given threshold length)
// in a squence of strings, then iterates through the keys to find the one that
// occurs the most frequently.

func FrequencyCounter(st searching.SymbolTable, path string, minLen int) *FrequencyResult {
	// compute frequency counts
	distinct, words := 0, 0
	in := testutil.NewInReadWords(path)

	// timer := util.NewStopwatch()

	for !in.IsEmpty() {
		word := in.ReadString()
		if len(word) < minLen {
			continue
		}
		words++
		key := searching.StringKey(word)
		if st.Contains(key) {
			value := st.Get(key).(int) + 1
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

	// log.Printf("%T, elapsed time: %.2f seconds", st, timer.ElapsedTime())

	return &FrequencyResult{max, st.Get(max), minLen, words, distinct}
}

func FrequencyCounterOrdered(st searching.OrderedSymbolTable, path string, minLen int) *FrequencyResult {
	// compute frequency counts
	distinct, words := 0, 0
	in := testutil.NewInReadWords(path)

	// timer := util.NewStopwatch()

	for !in.IsEmpty() {
		word := in.ReadString()
		if len(word) < minLen {
			continue
		}
		words++
		key := searching.StringKey(word)
		if st.Contains(key) {
			value := st.Get(key).(int) + 1
			st.Put(key, value)
		} else {
			st.Put(key, 1)
			distinct++
		}
	}

	// find a word with the highest frequency count
	var max searching.OSTKey = searching.StringKey("")
	st.Put(max, 0)
	for _, w := range st.Keys() {
		wInt := st.Get(w).(int)
		maxInt := st.Get(max).(int)
		if wInt > maxInt {
			max = w
		}
	}

	// log.Printf("%T, elapsed time: %.2f seconds", st, timer.ElapsedTime())

	return &FrequencyResult{max, st.Get(max), minLen, words, distinct}
}

func FrequencyCounterHash(st searching.HashSymbolTable, path string, minLen int) *FrequencyResult {
	// compute frequency counts
	distinct, words := 0, 0
	in := testutil.NewInReadWords(path)

	// timer := util.NewStopwatch()

	for !in.IsEmpty() {
		word := in.ReadString()
		if len(word) < minLen {
			continue
		}
		words++
		key := searching.StringHashKey(word)
		if st.Contains(key) {
			value := st.Get(key).(int) + 1
			st.Put(key, value)
		} else {
			st.Put(key, 1)
			distinct++
		}
	}

	// find a word with the highest frequency count
	var max searching.HashSTKey = searching.StringHashKey("")
	st.Put(max, 0)
	for _, w := range st.Keys() {
		wInt := st.Get(w).(int)
		maxInt := st.Get(max).(int)
		if wInt > maxInt {
			max = w
		}
	}

	// log.Printf("%T, elapsed time: %.2f seconds", st, timer.ElapsedTime())

	return &FrequencyResult{max, st.Get(max), minLen, words, distinct}
}

type FrequencyResult struct {
	max                     searching.STKey
	frequency               searching.STValue
	minLen, words, distinct int
}

func (r FrequencyResult) String() string {
	const str = "high-frequency word: %v, frequency: %v, minLen: %d, total words: %d, distinct words: %d"
	return fmt.Sprintf(str, r.max, r.frequency, r.minLen, r.words, r.distinct)
}

const (
	tinyTalePath = "testdata/tinyTale.txt"
	talePath     = "testdata/tale.txt.gz"
)

func ExampleSequentialSearchST_frequency() {
	st := searching.NewSequentialSearchST()

	// elapsed time: 0.00 seconds
	result := FrequencyCounter(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// result = FrequencyCounter(st, "https://algs4.cs.princeton.edu/31elementary/tale.txt", 8)
	// elapsed time: 2.57 seconds
	result = FrequencyCounter(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: of, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128
}

func ExampleBinarySearchST_frequency() {
	st := searching.NewBinarySearchST()

	// elapsed time: 0.00 seconds
	result := FrequencyCounterOrdered(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// elapsed time: 0.16 seconds
	result = FrequencyCounterOrdered(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: it, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128
}

func ExampleBST_frequency() {
	st := searching.NewBST()

	// elapsed time: 0.00 seconds
	result := FrequencyCounterOrdered(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// elapsed time: 0.11 seconds
	result = FrequencyCounterOrdered(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: it, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128
}

func ExampleRedBlackBST_frequency() {
	st := searching.NewRedBlackBST()

	// elapsed time: 0.00 seconds
	result := FrequencyCounterOrdered(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// elapsed time: 0.07 seconds
	result = FrequencyCounterOrdered(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: it, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128
}

func ExampleSeparateChainingHashST_frequency() {
	st := searching.NewSeparateChainingHashST()

	// elapsed time: 0.00 seconds
	result := FrequencyCounterHash(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// elapsed time: 0.10 seconds
	result = FrequencyCounterHash(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: the, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128
}

func ExampleLinearProbingHashST_frequency() {
	st := searching.NewLinearProbingHashST()

	// elapsed time: 0.00 seconds
	result := FrequencyCounterHash(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// elapsed time: 0.09 seconds
	result = FrequencyCounterHash(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: of, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128
}
