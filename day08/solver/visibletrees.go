package solver

import (
	"strconv"

	"github.com/Zairian/advent-of-code-2022/utils"
)

var treeMap = parseInput()

func CountVisible() int {
	counter := 0

	for i := range treeMap {
		for l := range treeMap[i] {
			if checkVisible(l, i, treeMap) {
				counter++
			}
		}
	}
	return counter
}

func BestScenicScore() int {
	bestScore := 0

	for i := range treeMap {
		for l := range treeMap[i] {
			score := calcScenicScore(l, i, treeMap)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func parseInput() [][]int {
	var treeMap [][]int
	s := utils.GetScanner("./day08/input.txt")
	defer utils.CloseOS()

	for i := 0; s.Scan(); i++ {
		treeMap = append(treeMap, make([]int, 0))
		for _, char := range s.Text() {
			num, _ := strconv.Atoi(string(char))
			treeMap[i] = append(treeMap[i], num)
		}
	}
	return treeMap
}

func checkVisible(x, y int, treeMap [][]int) bool {
	if x == 0 || y == 0 || x == len(treeMap[y])-1 || y == len(treeMap)-1 {
		return true
	}

	for i := x - 1; i >= 0; i-- {
		if treeMap[y][x] <= treeMap[y][i] {
			break
		}
		if treeMap[y][x] > treeMap[y][i] && i == 0 {
			return true
		}
	}
	for i := x + 1; i < len(treeMap[y]); i++ {
		if treeMap[y][x] <= treeMap[y][i] {
			break
		}
		if treeMap[y][x] > treeMap[y][i] && i == len(treeMap[y])-1 {
			return true
		}
	}
	for i := y - 1; i >= 0; i-- {
		if treeMap[y][x] <= treeMap[i][x] {
			break
		}
		if treeMap[y][x] > treeMap[i][x] && i == 0 {
			return true
		}
	}
	for i := y + 1; i < len(treeMap); i++ {
		if treeMap[y][x] <= treeMap[i][x] {
			break
		}
		if treeMap[y][x] > treeMap[i][x] && i == len(treeMap)-1 {
			return true
		}
	}
	return false
}

func calcScenicScore(x, y int, treeMap [][]int) int {
	dirScore := 0
	totScore := 1
	for i := x - 1; i >= 0; i-- {
		dirScore++
		if treeMap[y][x] <= treeMap[y][i] {
			break
		}
	}
	if dirScore != 0 {
		totScore *= dirScore
		dirScore = 0
	}
	for i := x + 1; i < len(treeMap[y]); i++ {
		dirScore++
		if treeMap[y][x] <= treeMap[y][i] {
			break
		}
	}
	if dirScore != 0 {
		totScore *= dirScore
		dirScore = 0
	}
	for i := y - 1; i >= 0; i-- {
		dirScore++
		if treeMap[y][x] <= treeMap[i][x] {
			break
		}
	}
	if dirScore != 0 {
		totScore *= dirScore
		dirScore = 0
	}
	for i := y + 1; i < len(treeMap); i++ {
		dirScore++
		if treeMap[y][x] <= treeMap[i][x] {
			break
		}
	}
	if dirScore != 0 {
		totScore *= dirScore
		dirScore = 0
	}
	return totScore
}
