package client_test

import (
	"fmt"

	// "github.com/youngzhu/algs4-go/searching"
	"github.com/youngzhu/algs4-go/searching/client"
)

func ExampleSequentialSearchST() {
	result := client.FrequencyCounter("testdata/tinyTale.txt", 1)
	fmt.Println(result.String())

	// fmt.Println()

	result = client.FrequencyCounter("testdata/tale.txt.gz", 8)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: of, frequency: 10, minLen: 1, total words: 60, distinct words: 20
	// high-frequency word: business, frequency: 122, minLen: 8, total words: 14350, distinct words: 5131

}