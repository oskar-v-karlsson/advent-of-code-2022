package solver

import "fmt"

func Solve1() {
	fmt.Printf("The start of the packet marker is at character: %v\n", GetStartOf(4))
}

func Solve2() {
	fmt.Printf("The start of the message marker is at character: %v\n", GetStartOf(14))
}
