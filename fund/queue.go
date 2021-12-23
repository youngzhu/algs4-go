package fund

import "strings"

// FIFO Queue.
// A FIFO queue is a collection that is based on the first-in-first-out (FIFO)
// policy. The pilicy of doing task in the same order that they arrive is one
// that we encounter frequently in everyday life: from people waiting in line
// at a theater, to cars waiting in line at a toll booth, to tasks waiting to 
// be serviced by an application on your computer.

// Queue implemented using a linked list
type Queue struct {
	first *Node // beginning of queue
	last *Node // end of queue
	n int // number of elements on queue
}

func NewQueue() *Queue {
	return &Queue{}
}

// Returns the item least recently added to the queue
func (q *Queue) Peek() Item {
	if q.IsEmpty() {
		panic("This queue is empty")
	}
	return q.first.item
}

// Adds the item to the queue
func (q *Queue) Enqueue(item Item) {
	oldLast := q.last
	q.last = newNode(item, nil)

	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.n++
}

// Removes and returns the item on this queue that was least recently added
func (q *Queue) Dequeue() Item {
	if q.IsEmpty() {
		panic("This queue is empty")
	}
	item := q.first.item
	q.first = q.first.next
	q.n--
	return item
}

// Returns the number of items in this queue
func (q *Queue) Size() int {
	return q.n
}

// Return true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Iterator() Iterator {
	items := make([]interface{}, q.n)

	i := 0
	cur := q.first
	for i < q.n {
		items[i] = cur.item
		cur = cur.next
		i++
	}

	return Iterator(items)
}

func (q *Queue) String() string {
	var ss []string

	for _, v := range q.Iterator() {
		ss = append(ss, v.(string))
	}

	return "[" + strings.Join(ss, ", ") + "]"
}