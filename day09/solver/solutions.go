package solver

import "fmt"

func Solve() {
	fmt.Printf("The short rope's tail visits: %v coordinates\n", SumVisited())
	fmt.Printf("The long rope's tail visits: %v coordinates\n", SumVisitedLong())
}
