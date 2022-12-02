package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,

	"X": 1,
	"Y": 2,
	"Z": 3,
}

func CalculateScore() int {
	f, err := os.Open("input.txt")
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
