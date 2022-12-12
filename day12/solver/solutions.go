package solver

import "fmt"

func Solve() {
	v, r := calcShortPath()
	fmt.Printf("Shortest path starting from one start point is: %v with the route %s\n\n", v, r)
	v, r = calcMultiStartShortPath()
	fmt.Printf("Shortest path starting from multiple starting points is: %v with the route %s", v, r)
}
