package fund

import (
	"strings"
	"fmt"
)

// Bags.
// A bag is a collection where removing items is not supported -- its purpose
// is to provide clients with the ability to collect items and than to iterate
// through the collected items.

// implemented using a singly linked list
type Bag struct {
	first *Node // beginning of bag
	n     int   // number of elements in bag
}

func NewBag() *Bag {
	return &Bag{}
}

// Adds the item to this bag
func (b *Bag) Add(item Item) {
	b.first = newNode(item, b.first)
	b.n++
}

func (b *Bag) Iterator() Iterator {
	items := make([]interface{}, b.n)

	for i, cur := 0, b.first; i < b.n; i, cur = i+1, cur.next {
		items[i] = cur.item
	}

	return Iterator(items)
}

// Returns true if this bag is empty
func (b *Bag) IsEmpty() bool {
	return b.first == nil
}

// Returns the number of items in this bag
func (b *Bag) Size() int {
	return b.n
}

func (b *Bag) String() string {
	var ss []string

	for _, v := range b.Iterator() {
		ss = append(ss, fmt.Sprint(v))
	}

	return "[" + strings.Join(ss, ", ") + "]"
}