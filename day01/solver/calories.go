package solver

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/Zairian/advent-of-code-2022/utils"
)

// Counts calories of individual elves and returns the calory count of the elf carrying the most
func calorieCount() int {
	s := utils.GetScanner("./day01/input.txt")
	defer utils.CloseOS()

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

// Count calories of individual elves and returns the sum of the largest three
func calorieCountThree() int {
	f, err := os.Open("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var sortedCalories []int
	sum := 0

	for s.Scan() {
		if s.Text() == "" {
			sortedCalories = append(sortedCalories, sum)
			sort.Sort(sort.Reverse(sort.IntSlice(sortedCalories)))
			sum = 0
			continue
		}

		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalln(err)
		}
		sum += i
	}

	largestThree := 0

	for i := 0; i < 3; i++ {
		largestThree += sortedCalories[i]
	}
	return largestThree
}
