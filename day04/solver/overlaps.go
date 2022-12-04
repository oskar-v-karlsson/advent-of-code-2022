package solver

import (
	"log"
	"strconv"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

func sumFullOverlaps() int {
	s := utils.GetScanner("input.txt")
	defer utils.CloseOS()

	sum := 0

	for s.Scan() {
		boundA, boundB := getPairBounds(strings.Split(s.Text(), ","))
		if checkFullOverlap(boundA, boundB) {
			sum += 1
		}
	}
	return sum
}

func sumPartialOverlaps() int {
	s := utils.GetScanner("input.txt")
	defer utils.CloseOS()

	sum := 0

	for s.Scan() {
		boundA, boundB := getPairBounds(strings.Split(s.Text(), ","))
		if checkPartialOverlap(boundA, boundB) {
			sum += 1
		}
	}
	return sum
}

// Returns the bounds of each range
func getPairBounds(input []string) ([]int, []int) {
	boundA := make([]int, 2)
	boundB := make([]int, 2)

	boundA[0], boundA[1] = getBounds(input[0])

	boundB[0], boundB[1] = getBounds(input[1])

	return boundA, boundB
}

// Returns the lower and upper bound from a string with the format X-Y
func getBounds(s string) (int, int) {
	bounds := strings.Split(s, "-")
	lowerBound, err := strconv.Atoi(bounds[0])
	if err != nil {
		log.Fatal(err)
	}

	upperBound, err := strconv.Atoi(bounds[1])
	if err != nil {
		log.Fatal(err)
	}

	return lowerBound, upperBound
}

// Returns true if one of the ranges fully overlap the other
func checkFullOverlap(boundA, boundB []int) bool {
	switch {
	case boundA[0] >= boundB[0] && boundA[1] <= boundB[1]:
		return true
	case boundB[0] >= boundA[0] && boundB[1] <= boundA[1]:
		return true
	default:
		return false
	}
}

// Returns true if one of the ranges partially overlap the other
func checkPartialOverlap(boundA, boundB []int) bool {
	return boundA[0] <= boundB[1] && boundA[1] >= boundB[0]
}
