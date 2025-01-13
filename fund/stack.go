package fund

import (
	"fmt"
	"strings"
)

// Pushdown stack.
// A pushdown stack is a collection that is based on the last-in-first-out (LIFO)
// policy. When you click a hyperlink, your browser displays the new page (and
// pushes onto a stack). You can keep clicking on hyperlinks to visit new pages,
// but you can always revisit the previous page by clicking the back button
// (popping it from the stack).

// Stack implemented using a singly linked list
type Stack struct {
	top  *Node // top of stack
	size int   // size of the stack
}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds the item to this stack
func (s *Stack) Push(item Item) {
	s.top = newNode(item, s.top)
	s.size++
}

// Pop removes and returns the item most recently added to this stack
func (s *Stack) Pop() Item {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	item := s.top.item // save item to return
	s.top = s.top.next // delete first node
	s.size--

	return item
}

// Peek returns (but does not remove) the item most recently added to this stack
func (s *Stack) Peek() Item {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	return s.top.item
}

// IsEmpty returns true if this stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of items in this stack
func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Iterator() Iterator {
	items := make([]interface{}, s.size)

	i := 0
	cur := s.top
	for i < s.size {
		items[i] = cur.item
		cur = cur.next
		i++
	}

	return items
}

func (s *Stack) String() string {
	var ss []string

	for _, v := range s.Iterator() {
		ss = append(ss, fmt.Sprint(v))
	}

	return "[" + strings.Join(ss, ", ") + "]"
}
