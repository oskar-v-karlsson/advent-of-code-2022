package utils

import (
	"log"
	"strconv"
	"strings"
)

func SumOverlaps() int {
	s := getScanner("input.txt")
	defer closeOS()

	sum := 0

	for s.Scan() {
		pair := expandPair(strings.Split(s.Text(), ","))
		if checkOverlap(pair) {
			sum += 1
		}
	}
	return sum
}

// Takes a string slice with the input format {X-Y, Z-W} and returns a string slice with each range expanded
func expandPair(input []string) []string {
	expandedPair := make([]string, 2)

	lowerBoundA, upperBoundA := getBounds(input[0])
	expandedPair[0] = convRangeToString(lowerBoundA, upperBoundA)

	lowerBoundB, upperBoundB := getBounds(input[1])
	expandedPair[1] = convRangeToString(lowerBoundB, upperBoundB)

	return expandedPair
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

// Converts a lower bound and an upper bound to a string eg. "3" "4" "5" "6" "7"
func convRangeToString(lowerBound, upperBound int) string {
	s := ""
	for i := lowerBound; i <= upperBound; i++ {
		s += "\""
		s += strconv.Itoa(i)
		s += "\" "
	}
	return s
}

// Returns true if one of the ranges fully overlap the other
func checkOverlap(pair []string) bool {
	if len([]rune(pair[0])) < len([]rune(pair[1])) {
		return strings.Contains(pair[1], pair[0])
	} else {
		return strings.Contains(pair[0], pair[1])
	}
}
