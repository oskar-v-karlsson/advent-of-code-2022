package main

import (
	"fmt"

	"github.com/Zairian/advent-of-code-2022/day03/utils"
)

func main() {
	fmt.Printf("The sum of the priorities is: %v\n", utils.SumIdenticalPriorities())
	fmt.Printf("The summed priorities of the groups are: %v\n", utils.SumGroupPriority())
}
