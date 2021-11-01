package client_test

import (
	"fmt"

	// "github.com/youngzhu/algs4-go/searching"
	"github.com/youngzhu/algs4-go/searching/client"
)

func ExampleSequentialSearchST() {
	result := client.FrequencyCounter("testdata/tinyTale.txt", 1)
	fmt.Println(result.String())

	// Output:
	// high-frequency word: of, frequency: 10
	// minLen: 1, total words: 60, distinct words: 20

}