package solver

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

func crateMover9000() string {
	crates := make([][]string, 9)

	s := utils.GetScanner("./day05/input.txt")
	defer utils.CloseOS()

	parseCrates(s, crates)

	for s.Scan() {
		if s.Text() != "" {
			number, from, to := parseActions(s.Text())
			moveCrates(number, from, to, crates)
		}
	}
	return getTopCrates(crates)
}

func crateMover9001() string {
	crates := make([][]string, 9)

	s := utils.GetScanner("./day05/input.txt")
	defer utils.CloseOS()

	parseCrates(s, crates)

	for s.Scan() {
		if s.Text() != "" {
			number, from, to := parseActions(s.Text())
			moveMultipleCrates(number, from, to, crates)
		}
	}
	return getTopCrates(crates)
}

// Returns initial setup of crates as a 2d matrix
func parseCrates(s *bufio.Scanner, crates [][]string) {
	for i := 0; i < len(crates); i++ {
		s.Scan()
		crates[i] = strings.Split(s.Text(), ",")
	}
}

// Returns the number and the indexes from where crates should be moved from and the index where it should be moved to
func parseActions(action string) (int, int, int) {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(action, -1)

	number, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])

	return number, from - 1, to - 1
}

// Move crates one at a time
func moveCrates(number, from, to int, crates [][]string) {
	for i := 0; i < number; i++ {
		topCrateIndex := len(crates[from]) - 1
		crates[to] = append(crates[to], crates[from][topCrateIndex])
		// Remove moved crate
		crates[from] = crates[from][:len(crates[from])-1]
	}
}

// Move multiple crates while keeping order
func moveMultipleCrates(number, from, to int, crates [][]string) {
	for i := number; i > 0; i-- {
		crateToBeMoved := len(crates[from]) - i
		crates[to] = append(crates[to], crates[from][crateToBeMoved])
	}
	// Remove moved crates
	crates[from] = crates[from][:len(crates[from])-number]
}

func getTopCrates(crates [][]string) string {
	topCrates := ""
	for _, stack := range crates {
		topCrates += stack[len(stack)-1]
	}
	return topCrates
}
