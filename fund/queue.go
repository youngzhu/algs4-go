package fund

import (
	"fmt"
	"strings"
)

// FIFO Queue.
// A FIFO queue is a collection that is based on the first-in-first-out (FIFO)
// policy. The policy of doing task in the same order that they arrive is one
// that we encounter frequently in everyday life: from people waiting in line
// at a theater, to cars waiting in line at a toll booth, to tasks waiting to
// be serviced by an application on your computer.

// Queue implemented using a linked list
type Queue struct {
	first *Node // beginning of queue
	last  *Node // end of queue
	size  int   // number of elements on queue
}

func NewQueue() *Queue {
	return &Queue{}
}

// Peek returns the item least recently added to the queue
func (q *Queue) Peek() Item {
	if q.IsEmpty() {
		panic("This queue is empty")
	}
	return q.first.item
}

// Enqueue adds the item to the queue
func (q *Queue) Enqueue(item Item) {
	oldLast := q.last
	q.last = newNode(item, nil)

	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.size++
}

// Dequeue removes and returns the item on this queue that was least recently added
func (q *Queue) Dequeue() Item {
	if q.IsEmpty() {
		panic("This queue is empty")
	}
	item := q.first.item
	q.first = q.first.next
	q.size--
	return item
}

// Size returns the number of items in this queue
func (q *Queue) Size() int {
	return q.size
}

// IsEmpty return true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Iterator() Iterator {
	items := make([]interface{}, q.size)

	i := 0
	cur := q.first
	for i < q.size {
		items[i] = cur.item
		cur = cur.next
		i++
	}

	return items
}

func (q *Queue) String() string {
	var ss []string

	for _, v := range q.Iterator() {
		ss = append(ss, fmt.Sprint(v))
	}

	return "[" + strings.Join(ss, ", ") + "]"
}
