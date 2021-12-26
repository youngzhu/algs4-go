package graphs

// Search find vertices connected to a source vertex s
type Search interface {
	Marked(v int) bool // Is v connected to s?
	Count() int        // How many vertices are connected to s?
}
