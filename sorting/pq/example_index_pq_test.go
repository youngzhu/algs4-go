package pq_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/sorting/pq"
	"github.com/youngzhu/algs4-go/testutil"
)

// A client that reads sorted text files,
// merge them together into a sorted output

func ExampleMinIndexPQ() {
	streams := []testutil.In{
		*testutil.NewInReadWords("testdata/m1.txt"),
		*testutil.NewInReadWords("testdata/m2.txt"),
		*testutil.NewInReadWords("testdata/m3.txt"),
	}

	merge(pq.NewMinIndexPQ(len(streams)), streams)

	// Output:
	// A A B B B C D E F F G H I I J N P Q Q Z
}

// merge the sorted input streams
func merge(ipq pq.IndexPQ, streams []testutil.In) {
	n := len(streams)
	for i := 0; i < n; i++ {
		if streams[i].HasNext() {
			item := pq.StringItem(streams[i].ReadString())
			ipq.Insert(i, item)
		}
	}

	// extract and print min and read next from its stream
	for !ipq.IsEmpty() {
		fmt.Print(ipq.HighestPriorityItem(), " ")
		i := ipq.Delete()
		// fmt.Println(ipq)
		if streams[i].HasNext() {
			item := pq.StringItem(streams[i].ReadString())
			ipq.Insert(i, item)
		}
	}
}
