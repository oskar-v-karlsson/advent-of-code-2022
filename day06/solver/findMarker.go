package solver

import (
	"github.com/Zairian/advent-of-code-2022/utils"
)

func GetStartOf(distinct int) int {
	// TODO Change path
	s := utils.GetScanner("./day06/input.txt")

	s.Scan()

	msg := s.Text()

	// Fill initial array with fourth rune in the msg sequence
	var prevFour []rune

	for i, char := range msg {
		prevFour = append(prevFour, char)
		// Remove first element
		if len(prevFour) > distinct {
			prevFour = prevFour[1:]
		}
		if !checkDuplicates(prevFour) && len(prevFour) == distinct {
			return i + 1
		}
	}
	return 0
}

func checkDuplicates(dupArr []rune) bool {
	duplicates := false
	for i := range dupArr {
		for j := i + 1; j < len(dupArr); j++ {
			if dupArr[i] == dupArr[j] {
				duplicates = true
				break
			}
		}
		if duplicates == true {
			return duplicates
		}
	}
	return duplicates
}
