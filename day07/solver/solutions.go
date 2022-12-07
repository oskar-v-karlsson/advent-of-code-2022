package solver

import "fmt"

func Solve() {
	solution, size := CalcSmallDirectories()
	fmt.Printf("The summed size of the directories smaller than 100000 is: %v\n", solution)
	fmt.Printf("The smallest deletable directory's size is: %v", CalcDeletable(size))
}
