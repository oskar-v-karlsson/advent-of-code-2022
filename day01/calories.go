package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Counts calories of individual elves and returns the calory count of the elf carrying the most
func CalorieCount() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	largest := 0
	sum := 0

	for s.Scan() {
		if s.Text() == "" {
			if sum > largest {
				largest = sum
			}
			sum = 0
			continue
		}

		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalln(err)
		}
		sum += i
	}

	return largest
}
