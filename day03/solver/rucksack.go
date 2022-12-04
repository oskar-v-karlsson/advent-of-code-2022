package solver

import (
	"bytes"
	"unicode"

	"github.com/Zairian/advent-of-code-2022/day03/utils"
)

func sumIdenticalPriorities() int {
	s := utils.GetScanner()
	defer utils.CloseOS()

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

func sumGroupPriority() int {
	s := utils.GetScanner()
	defer utils.CloseOS()

	sum := 0
	iterator := 0

	group := make([]string, 3)

	for s.Scan() {
		group[iterator] = s.Text()
		if iterator == 2 {
			sum += groupPriority(group)
			iterator = 0
			continue
		}
		iterator++
	}

	return sum
}

// Returns the priority of the identical badge in the group
func groupPriority(group []string) int {
	for _, char := range group[0] {
		if bytes.ContainsRune([]byte(group[1]), char) && bytes.ContainsRune([]byte(group[2]), char) {
			if unicode.IsLower(char) {
				return (int(char) - 96)
			}
			return (int(char) - 38)
		}
	}

	return 0
}
