package fund

// Linked List.
// A linked list is a recursive data structure that is either empty (nil) or 
// a reference to a node having a item and a reference to a linked list.

type Item interface{}

type Node struct {
	item Item
	next *Node
}

func newNode(item Item, next *Node) *Node {
	return &Node{item, next}
}