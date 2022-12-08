package solver

import "fmt"

func Solve1() {
	fmt.Printf("Amount of trees that are visible from any angle: %v\n", CountVisible())
}

func Solve2() {
	fmt.Printf("Best scenic score of all trees is: %v", BestScenicScore())
}
