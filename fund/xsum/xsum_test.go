package xsum_test

import (
	"log"
	"os"
	"testing"

	. "github.com/youngzhu/algs4-go/fund/xsum"
	"github.com/youngzhu/algs4-go/testutil"
)

func readInts(path string) []int {
	in := testutil.NewInReadWords(path)
	return in.ReadAllInts()
}

var ints1K = readInts("testdata/1Kints.txt")
var ints2K = readInts("testdata/2Kints.txt.gz")
var ints4K = readInts("testdata/4Kints.txt.gz")

func TestMain(m *testing.M) {
	log.Println("before..")

	r := m.Run()

	log.Println("after...")

	os.Exit(r)
}

func TestThreeSumCount(t *testing.T) {
	// t.Parallel()

	want := []int{70, 528, 4039}

	got := []int{
		ThreeSumCount(ints1K),
		ThreeSumCount(ints2K),
		ThreeSumCount(ints4K),
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("got: %v; want: %v", got, want)
		}
	}
}

func TestThreeSumCountFast(t *testing.T) {
	// t.Parallel()

	want := []int{70, 528, 4039}

	got := []int{
		ThreeSumCountFast(ints1K),
		ThreeSumCountFast(ints2K),
		ThreeSumCountFast(ints4K),
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("got: %v; want: %v", got, want)
		}
	}
}

func TestTwoSumCount(t *testing.T) {
	// t.Parallel()

	want := []int{1, 2, 3}

	got := []int{
		TwoSumCount(ints1K),
		TwoSumCount(ints2K),
		TwoSumCount(ints4K),
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("got: %v; want: %v", got, want)
		}
	}

}

func TestTwoSumCountFast(t *testing.T) {
	// t.Parallel()

	want := []int{1, 2, 3}

	got := []int{
		TwoSumCountFast(ints1K),
		TwoSumCountFast(ints2K),
		TwoSumCountFast(ints4K),
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("got: %v; want: %v", got, want)
		}
	}

}

// go test -v
// TestThreeSumCount (12.18s)
// TestThreeSumCountFast (0.40s)
// TestTwoSumCount (0.01s)
// TestTwoSumCountFast (0.00s)
