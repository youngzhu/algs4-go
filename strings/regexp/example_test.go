package regexp_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/strings/regexp"
)

var testCases = []struct {
	regExp, text string
}{
	{"(A*B|AC)D", "AAAABD"},
	{"(A*B|AC)D", "AAAAC"},
	{"(a|(bc)*d)*", "abcbcd"},
	{"(a|(bc)*d)*", "abcbcbcdaaaabcbcdaaaddd"},
}

func ExampleNFA() {

	for _, tc := range testCases {
		nfa := regexp.NewNFA(tc.regExp)
		result := nfa.Recognizes(tc.text)
		fmt.Println(result)
	}

	// Output:
	// true
	// false
	// true
	// true
}
