package mst

// Proposition.
// Prim's algorithm computes the MST of any connected edge-weighted graph. 
// The lazy version of Prim's algorithm uses space proportional to E and time
// proportional to ElogE (in the worst case) to compute the MST of a connected
// edge-weighted graph with E edges and V vertices.
// The eager version uses space proportional to V and time proportional to 
// ElogV (in the worst case).