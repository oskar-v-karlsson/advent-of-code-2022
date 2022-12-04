package solver

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,

	"X": 1,
	"Y": 2,
	"Z": 3,
}

func calculateScore() int {
	s := utils.GetScanner("./day02/input.txt")
	defer utils.CloseOS()

	totalScore := 0

	for s.Scan() {
		hands := strings.Split(s.Text(), " ")

		// Converts hand characters to their respective scores
		handValues := []int{scores[hands[0]], scores[hands[1]]}

		totalScore += checkResults(handValues)
		totalScore += handValues[1]
	}
	return totalScore
}

// Returns 0 if loss, 3 if draw and 6 if win
func checkResults(h []int) int {
	comparison := h[1] - h[0]
	switch {
	case comparison == 1 || comparison == -2:
		return 6
	case comparison == 0:
		return 3
	default:
		return 0
	}
}

func calculateTrueScore() int {
	f, err := os.Open("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	totalScore := 0

	for s.Scan() {
		hands := strings.Split(s.Text(), " ")

		// Converts hand characters to their respective scores
		handValues := []int{scores[hands[0]], scores[hands[1]]}

		switch hands[1] {
		case "X":
			totalScore += createLossHand(handValues)
		case "Y":
			totalScore += createDrawHand(handValues)
			totalScore += 3
		case "Z":
			totalScore += createWinHand(handValues)
			totalScore += 6
		}
	}
	return totalScore
}

// Returns the number corresponding to winning hand
func createWinHand(h []int) int {
	if h[0] == 3 {
		return 1
	}
	return h[0] + 1
}

// Returns the number corresponding to draw hand
func createDrawHand(h []int) int {
	return h[0]
}

// Returns the number corresponding to losing hand
func createLossHand(h []int) int {
	if h[0] == 1 {
		return 3
	}
	return h[0] - 1
}
