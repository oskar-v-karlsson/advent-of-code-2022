package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

type ropeEnd struct {
	x int
	y int
}

func SumVisited() int {
	counter := 1
	visited := map[string]bool{"0, 0": true}
	var h, t ropeEnd

	s := utils.GetScanner("input.txt")
	defer utils.CloseOS()

	for s.Scan() {
		dir, steps := parseCMD(s.Text())
		for i := 0; i < steps; i++ {
			moveHead(dir, &h)
			moveTail(&h, &t)
			if checkVisited(t, visited) {
				counter++
			}
		}
	}
	return counter
}

func SumVisitedLong() int {
	counter := 1
	visited := map[string]bool{"0, 0": true}
	var h ropeEnd
	t := make([]ropeEnd, 9)

	s := utils.GetScanner("input.txt")
	defer utils.CloseOS()

	for s.Scan() {
		dir, steps := parseCMD(s.Text())
		for i := 0; i < steps; i++ {
			moveHead(dir, &h)
			for i := range t {
				if i == 0 {
					moveTail(&h, &t[0])
				} else {
					moveTail(&t[i-1], &t[i])
				}
			}
			if checkVisited(t[8], visited) {
				counter++
			}
		}
	}
	return counter
}

// Returns the direction of movement and steps
func parseCMD(input string) (string, int) {
	cmd := strings.Split(input, " ")
	dir := cmd[0]
	steps, _ := strconv.Atoi(cmd[1])

	return dir, steps
}

func moveHead(dir string, head *ropeEnd) {
	switch dir {
	case "U":
		head.y -= 1
	case "R":
		head.x += 1
	case "D":
		head.y += 1
	case "L":
		head.x -= 1
	}
}

// Moves the tail according to rules and returns true if the moved position hasn't been visited
func moveTail(head, tail *ropeEnd) {
	if !checkNeighbour(*head, *tail) {
		if head.x == tail.x {
			tail.y = tail.y + (head.y-tail.y)/utils.IntAbs(head.y-tail.y)
		} else if head.y == tail.y {
			tail.x = tail.x + (head.x-tail.x)/utils.IntAbs(head.x-tail.x)
		} else {
			tail.x = tail.x + (head.x-tail.x)/utils.IntAbs(head.x-tail.x)
			tail.y = tail.y + (head.y-tail.y)/utils.IntAbs(head.y-tail.y)
		}
	}
}

// Return true if the supplied points are neighbours
func checkNeighbour(head ropeEnd, tail ropeEnd) bool {
	distanceX := (head.x - tail.x)
	distanceY := (head.y - tail.y)
	if distanceX > 1 || distanceY > 1 || distanceX < -1 || distanceY < -1 {
		return false
	}
	return true
}

// Return true if this coordinate have never been visited
func checkVisited(tail ropeEnd, visited map[string]bool) bool {
	key := fmt.Sprintf("%v, %v", tail.x, tail.y)
	if !visited[key] {
		visited[key] = true
		return true
	}
	return false
}
