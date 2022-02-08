package suffix_test

import (
	"fmt"
	"github.com/youngzhu/algs4-go/context/suffix"
	"github.com/youngzhu/algs4-go/util"
	"regexp"
)

func ExampleSuffixArray() {
	in := util.NewIn("testdata/abra.txt")
	text := in.ReadAll()

	sa := suffix.NewSuffixArray(text)

	testSuffixArray(sa)

	// Output:
	//----------------------------
	//i  idx lcp rnk ith
	//----------------------------
	//0  11  -   0   !
	//1  10  0   1   A!
	//2  7   1   2   ABRA!
	//3  0   4   3   ABRACADABRA!
	//4  3   1   4   ACADABRA!
	//5  5   1   5   ADABRA!
	//6  8   0   6   BRA!
	//7  1   3   7   BRACADABRA!
	//8  4   0   8   CADABRA!
	//9  6   0   9   DABRA!
	//10 9   0   10  RA!
	//11 2   2   11  RACADABRA!
}

func testSuffixArray(sa suffix.SuffixArrayInterface) {
	fmt.Println("----------------------------")
	fmt.Println("i  idx lcp rnk ith")
	fmt.Println("----------------------------")

	for i := 0; i < sa.Length(); i++ {
		idx := sa.Index(i)
		ith := sa.Select(i)
		rank := sa.Rank(ith)
		if i == 0 {
			fmt.Printf("%-3d%-3d %-3s %-3d %s\n", i, idx, "-", rank, ith)
		} else {
			lcp := sa.LCP(i)
			fmt.Printf("%-3d%-3d %-3d %-3d %s\n", i, idx, lcp, rank, ith)
		}
	}
}

// Longest Repeated Substring
func ExampleLongestRepeatedSubstring() {
	lrs := suffix.LongestRepeatedSubstring("aaaaaaaa")
	fmt.Println(lrs)

	lrs = suffix.LongestRepeatedSubstring("abcdefg")
	fmt.Println(lrs)

	// Output:
	// "aaaaaaa"
	// ""
}

var space = regexp.MustCompile("\\s+")

func whitespace(s string) string {
	// not work
	//txt := strings.ReplaceAll(text, "\\s+", " ")

	//space := regexp.MustCompile("\\s+")
	t := string(space.ReplaceAll([]byte(s), []byte(" ")))

	return t
}
func ExampleLongestRepeatedSubstring_file() {
	in := util.NewIn("testdata/tinyTale.txt")
	text := in.ReadAll()
	s := whitespace(text)
	lrs := suffix.LongestRepeatedSubstring(s)
	fmt.Println(lrs)

	in = util.NewIn("testdata/mobydick.txt.gz")
	text = in.ReadAll()
	t := whitespace(text)
	lrs = suffix.LongestRepeatedSubstring(t)
	fmt.Println(lrs)

	// Output:
	// "st of times it was the "
	// ",- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th"
}

func ExampleLongestCommonSubstring() {
	in1 := util.NewIn("testdata/tale.txt.gz")
	text1 := in1.ReadAll()

	in2 := util.NewIn("testdata/mobydick.txt.gz")
	text2 := in2.ReadAll()

	s, t := whitespace(text1), whitespace(text2)
	lcs := suffix.LongestCommonSubstring(s, t)
	fmt.Println(lcs)

	// Output:
	// " seemed on the point of being "
}

func ExampleSuffixArrayX() {
	in := util.NewIn("testdata/abra.txt")
	text := in.ReadAll()

	sa := suffix.NewSuffixArrayX(text)
	testSuffixArray(sa)

	// Output:
	//----------------------------
	//i  idx lcp rnk ith
	//----------------------------
	//0  11  -   0   !
	//1  10  0   1   A!
	//2  7   1   2   ABRA!
	//3  0   4   3   ABRACADABRA!
	//4  3   1   4   ACADABRA!
	//5  5   1   5   ADABRA!
	//6  8   0   6   BRA!
	//7  1   3   7   BRACADABRA!
	//8  4   0   8   CADABRA!
	//9  6   0   9   DABRA!
	//10 9   0   10  RA!
	//11 2   2   11  RACADABRA!
}
