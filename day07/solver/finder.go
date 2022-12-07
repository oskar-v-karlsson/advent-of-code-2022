package solver

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

func CalcSmallDirectories() (int, int) {
	s := utils.GetScanner("input.txt")
	defer utils.CloseOS()

	sum, dirSize, _ := recursive(s, 0, 30000000, 30000000)

	return sum, dirSize
}

func CalcDeletable(dirSize int) int {
	s := utils.GetScanner("input.txt")
	defer utils.CloseOS()

	reqSpace := 30000000 - (70000000 - dirSize)

	_, _, deletable := recursive(s, 0, 3000000000, reqSpace)

	return deletable
}

func recursive(s *bufio.Scanner, sum int, smallest int, sizeNeeded int) (int, int, int) {
	totSum := sum
	dirSum := 0
	deletable := smallest
	s.Scan()

	// Sum file sizes in directory
	for !strings.Contains(s.Text(), "cd ") {
		if strings.Contains(s.Text(), "$") && strings.Contains(s.Text(), "dir") {
			continue
		}
		r := regexp.MustCompile(`[0-9]*`)
		num := r.FindString(s.Text())
		size, _ := strconv.Atoi(num)
		dirSum += size
		if !s.Scan() {
			return totSum, dirSum, deletable
		}
	}

	// Go up or down in directories
	for strings.Contains(s.Text(), "cd ") {
		switch {
		case s.Text() == "$ cd ..":
			if dirSum < deletable && dirSum >= sizeNeeded {
				deletable = dirSum
			}
			if dirSum > 100000 {
				return totSum, dirSum, deletable
			}
			return dirSum + totSum, dirSum, deletable
		case strings.Contains(s.Text(), "$ cd "):
			sum, subDirSize, small := recursive(s, 0, deletable, sizeNeeded)
			if small < deletable {
				deletable = small
			}
			totSum += sum
			dirSum += subDirSize
		}
		s.Scan()
	}

	return totSum, dirSum, deletable
}
