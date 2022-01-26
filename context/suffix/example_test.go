package suffix_test

import (
	"fmt"
	"github.com/youngzhu/algs4-go/context/suffix"
	"github.com/youngzhu/algs4-go/util"
)

func ExampleSuffixArray() {
	in := util.NewIn("testdata/abra.txt")
	text := in.ReadAll()

	fmt.Println("----------------------------")
	fmt.Println("i  idx lcp rnk ith")
	fmt.Println("----------------------------")

	sa := suffix.NewSuffixArray(text)
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
