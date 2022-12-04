package solver

import (
	"fmt"
)

func Solve1() {
	fmt.Printf("Amount of full overlaps: %v\n", sumFullOverlaps())
}

func Solve2() {
	fmt.Printf("Amount of partial overlaps: %v\n", sumPartialOverlaps())
}
