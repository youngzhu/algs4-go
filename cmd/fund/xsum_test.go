package main

import (
	"testing"
	"log"
	"os"

	"github.com/youngzhu/algs4-go/util"
)

func readInts(path string) []int {
	in := util.NewInReadWords(path)
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
	t.Parallel()

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