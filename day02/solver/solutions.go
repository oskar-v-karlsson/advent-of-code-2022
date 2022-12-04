package solver

import "fmt"

func Solve1() {
	fmt.Printf("Score without full strategy guide: %v points. \n", calculateScore())
}

func Solve2() {
	fmt.Printf("Score with full strategy guide: %v points. \n", calculateTrueScore())
}
