package sp

import (
	"fmt"
	"math"
	"log"

	"github.com/youngzhu/algs4-go/graphs/digraph"
	"github.com/youngzhu/algs4-go/util"
)

// Arbitrage detection.
// Consider a market for financial transations that is based on trading commodities.
// The table rates.txt shows conversion rates among currencies. The first line in
// the file is the number V of currencies; then the file has one line per currency,
// giving its name followed by the conversion rates to the other currencies. An
// arbitrage opportunity is a directed cycle such that the product of the exchange
// rates is greater than one. For example, our table syas that 1,000 U.S. dollars
// will buy 1,000 x 0.741 = 741 euros, then we can buy 741 x 1.366 = 1,012.206 
// Canadian dollars with our eruos, and finally, 1,012.206 x 0.995 = 1,007.14497
// U.S. dollars with our Canadian dollars, a 7.14497 dollar profit!
// 
// To formulate the arbitrage problem as a negative-cycle detection problem, 
// replace each weight by its logarithm, negated. With this change, computing
// path weights by multiplying edge weights in the original problem corresponds
// to adding them in the transformed problem.

// The routine identifies arbitrage opportunities in a currency-exchange network
// by solving the corresponding negative cycle detection problem.
func Arbitrage(path string) {
	in := util.NewInReadWords(path)

	n := in.ReadInt()
	name := make([]string, n)

	// create complete network
	g := digraph.NewEdgeWeightedDigraphN(n)
	for v := 0; v < n; v++ {
		name[v] = in.ReadString()
		for w := 0; w < n; w++ {
			rate := in.ReadFloat()
			e := digraph.NewDirectedEdge(v, w, -math.Log(rate))
			g.AddEdge(e)
		}
	}

	log.Println("c, stake")

	// find negative cycle
	spt := NewBellmanFordSP(*g, 0)
	if spt.HasNegativeCycle() {
		stake := 1000.0
		c := 0
		log.Println(c, stake, len(spt.NegativeCycle()))
		// for _, edge := range spt.NegativeCycle() {
		// 	c++
		// 	if c > 10 {
		// 		break
		// 	}
		// 	log.Println(c, stake)

		// 	e := edge.(digraph.DirectedEdge)
		// 	fmt.Printf("%10.5f %s ", stake, name[e.From()])
		// 	stake = stake * math.Exp(-e.Weight())
		// 	fmt.Printf("= %10.5f %s\n", stake, name[e.To()])
			
		// }
	} else {
		fmt.Println("No arbitrage opportunity")
	}
}