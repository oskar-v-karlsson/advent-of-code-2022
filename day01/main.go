package main

import (
	"fmt"

	"github.com/Zairian/advent-of-code-2022/utils"
)

func main() {
	fmt.Printf("Elf with largest amount of calories is carrying: %v calories!\n", utils.CalorieCount())
	fmt.Printf("The three elves with largest amount of calories are carrying: %v calories!\n", utils.CalorieCountThree())
}
