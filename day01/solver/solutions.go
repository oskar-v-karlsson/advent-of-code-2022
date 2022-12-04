package solver

import "fmt"

func Solve1() {
	fmt.Printf("Elf with largest amount of calories is carrying: %v calories!\n", calorieCount())
}

func Solve2() {
	fmt.Printf("The three elves with largest amount of calories are carrying: %v calories!\n", calorieCountThree())
}
