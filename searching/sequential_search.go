package searching

// Sequential search in an unordered linked list.
// SequentialSearchST implements a symbol table with a linked list of nodes
// that contains keys and values. To implement get(), we scan through the list,
// using equals() to compare the search key with the key in each node in th list.
// If we find the match, we return the associated value; if not, we return nil.
// To implement put(), we also scan through the list, using equals() to compare
// the client key with the key in each node in the list. If we find the match, 
// we update the value associated with that key to be the value given; if not,
// we create a new node with the given key and value, and insert it at the 
// beginning of the list. This method is known as sequential search.

type SequentialSearchST struct {
	n int // number of key-value pairs
	first *node // the linked list of key-value pairs
}

// a helper linked list data type
type node struct {
	key interface{}
	value interface{}
	next *node
}

func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{}
}

// Returns the value associated with the given key
// in this symbol table
func (st *SequentialSearchST) Get(key interface{}) interface{} {
	if key == nil {
		panic("argument to Get() is nil")
	}
	for x := st.first; x != nil; x = x.next {
		if x.key == key {
			return x.value
		}
	}
	return nil
}

// Inserts the specified key-value pair into the symbol table,
// overwriting the old value with the new value if the symbol
// table already contains the specified key.
// Deletes the specified key (and its associated value) from the
// symbol table if the specified value is nil
func (st *SequentialSearchST) Put(key, value interface{}) {
	if key == nil {
		panic("argument to Put() is nil")
	}
	if value == nil {
		st.Delete(key)
		return
	}

	for x := st.first; x != nil; x = x.next {
		if key == x.key {
			x.value = value
			return
		}
	}

	st.first = &node{key, value, st.first}
	st.n++
}

// Remove the specified key and its associated value from the symbol
// table (if the key is in the symbol table)
func (st *SequentialSearchST) Delete(key interface{}) {
	if key == nil {
		panic("argument to Delete() is nil")
	}
	st.first = st.deleteNode(st.first, key)
}

// delete key in linked list beginning at node x
// warning: function call stack too large if table is large
func (st *SequentialSearchST) deleteNode(x *node, key interface{}) *node {
	if x == nil {
		return nil
	}
	if key == x.key {
		st.n--
		return x.next
	}
	x.next = st.deleteNode(x.next, key)
	return x
}

// Return all keys in the symbol table
func (st *SequentialSearchST) Keys() []interface{} {
	var keys []interface{}
	for x := st.first; x != nil; x = x.next {
		keys = append(keys, x.key)
	}
	return keys
}