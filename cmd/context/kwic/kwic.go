package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/youngzhu/algs4-go/context/suffix"
	"github.com/youngzhu/algs4-go/testutil"
	"os"
	"regexp"
)

// KWIC: Key Word in Context
// Reads a string from a file specified as the first
// command-line argument; read an integer k specified
// as the second command-line argument; then repeatedly
// processes use queries, printing all occurrences of
// the given query string in the text string with k characters
// of surrounding context on either side.

var (
	file string
	k    int
)

func init() {
	flag.StringVar(&file, "f", "", "file name (path)")
	flag.IntVar(&k, "k", 0, "k characters of surrounding context")
}

// RUN
// go run kwic.go -f testdata/tale.txt.gz -k 10
func main() {
	flag.Parse()

	// read in text
	in, err := testutil.NewInWithError(file)
	if err != nil {
		panic(err)
	}
	text := in.ReadAll()
	space := regexp.MustCompile("\\s+")
	text = string(space.ReplaceAll([]byte(text), []byte(" ")))

	// build suffix array
	sa := suffix.NewSuffixArrayX(text)

	scanner := bufio.NewScanner(os.Stdin)
	// find all occurrences of queries and give context
	for {
		fmt.Print("Enter keyword:")
		scanner.Scan()
		keyword := scanner.Text()
		if keyword == "q" {
			os.Exit(0)
		}
		n := sa.Length()
		count := 0
		for i := sa.Rank(keyword); i < n; i++ {
			from1 := sa.Index(i)
			to1 := min(n, from1+len(keyword))
			if keyword != text[from1:to1] {
				break
			}
			from2 := max(0, from1-k)
			to2 := min(n, to1+k)
			fmt.Println(text[from2:to2])
			count++
		}
		fmt.Printf("==\nkeyword:\"%s\", occurrences:%d times", keyword, count)
		fmt.Println()
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
