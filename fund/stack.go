package fund

// Pushdown stack.
// A pushdown stack is a collection that is based on the last-in-first-out (LIFO)
// policy. When you click a hyperlink, your browser displays the new page (and
// pushes onto a stack). You can keep clicking on hyperlinks to visit new pages,
// but you can always revisit the previous page by clicking the back button
// (popping it from the stack).

// implemented using a singly linked list
type Stack struct {
	top *Node // top of stack
	n int // size of the stack
}

func NewStack() *Stack {
	return &Stack{}
}

// Adds the item to this stack
func (s *Stack) Push(item Item) {
	s.top = newNode(item, s.top)
	s.n++
}

// Removes and returns the item most recently added to this stack
func (s *Stack) Pop() Item {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	item := s.top.item // save item to return
	s.top = s.top.next // delete first node
	s.n--

	return item
}

// Returns (but does not remove) the item most recently added to this stack
func (s *Stack) Peek() Item {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	return s.top.item
}

// Returns true if this stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Returns the number of items in this stack
func (s *Stack) Size() int {
	return s.n
}

func (s *Stack) Iterator() Iterator {
	items := make([]interface{}, s.n)

	i := 0
	cur := s.top
	for i < s.n {
		items[i] = cur.item
		cur = cur.next
		i++
	}

	return Iterator(items)
}