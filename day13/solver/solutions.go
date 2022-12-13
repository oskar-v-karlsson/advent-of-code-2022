package solver

import "fmt"

func Solve() {
	solution1, solution2 := getCorrectOrder()
	fmt.Printf("The sum of the indices of correct packets is: %v\n", solution1)
	fmt.Printf("The decoder key is: %v", solution2)
}
