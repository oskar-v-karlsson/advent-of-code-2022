package main

import (
	"fmt"

	"github.com/Zairian/advent-of-code-2022/day02/utils"
)

func main() {
	fmt.Printf("Score without full strategy guide: %v points. \n", utils.CalculateScore())
	fmt.Printf("Score with full strategy guide: %v points. \n", utils.CalculateTrueScore())
}
