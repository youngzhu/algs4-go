package client_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/searching"
	"github.com/youngzhu/algs4-go/searching/client"
)

const (
	tinyTalePath = "testdata/tinyTale.txt"
	talePath     = "testdata/tale.txt.gz"
)

func ExampleSequentialSearchST() {
	st := searching.NewSequentialSearchST()
	result := client.FrequencyCounter(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// fmt.Println()

	result = client.FrequencyCounter(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: of, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128

}

func ExampleBinarySearchST() {
	st := searching.NewBinarySearchST()
	result := client.FrequencyCounter(st, tinyTalePath, 1)
	fmt.Println(result.String())

	// fmt.Println()

	result = client.FrequencyCounter(st, talePath, 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: it, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5128

}
