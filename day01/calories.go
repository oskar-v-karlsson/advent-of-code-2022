package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CalorieCount() {
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

	fmt.Println(largest)
}
