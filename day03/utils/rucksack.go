package utils

import (
	"bytes"
	"unicode"
)

func SumIdenticalPriorities() int {
	s := getScanner()
	defer closeOS()

	sum := 0

	for s.Scan() {
		sum += identicalPriority(s.Text())
	}
	return sum
}

// Returns the priority of the identical items in a rucksack
func identicalPriority(contents string) int {
	compartement1 := contents[0 : len(contents)/2]
	compartement2 := contents[len(contents)/2:]

	for _, char := range compartement1 {
		if bytes.ContainsRune([]byte(compartement2), char) {
			if unicode.IsLower(char) {
				return (int(char) - 96)
			}
			return (int(char) - 38)
		}
	}

	return 0
}
