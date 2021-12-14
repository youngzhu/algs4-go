package graphs

// Pairwise connections between items play a critical role in a vast array of
// computational applications. The relationships implied by these connections
// lead to a host of natural questions: Is there a way to connect one item to
// another by following the connections? How many other items are connected to
// a given item? What is the shortest chain of connections between this item
// and this other item?
// We progress through the four most important types of graph models:
// 1. Undirected Graphs (Graph) introduces the graph data type, including
// depth-first search and breadth-first search.
// 2. Directed Graphs (Digraph) introduces the digraph data type, including
// topological sort and strong components.
// 3. Minimum Spanning Trees (MST) describes the minimum spanning tree problem
// and two classic algorithms for solving it: Prim and Kruskal.
// 4. Shortest Paths (SP) introduces the shortest path problem and two classic
// algorithm for solving it: Dijkstra's algorithm and Bellman-Ford.

const InfinityDistance = 10000
