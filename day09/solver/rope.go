package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

type ropeEnd struct {
	x     int
	y     int
	prevX int
	prevY int
}

func SumVisited() int {
	counter := 1
	visited := make(map[string]bool)
	visited["0, 0"] = true
	h := ropeEnd{x: 0, y: 0}
	t := ropeEnd{x: 0, y: 0}

	s := utils.GetScanner("input.txt")
	defer utils.CloseOS()

	for s.Scan() {
		cmd := strings.Split(s.Text(), " ")
		dir := cmd[0]
		steps, _ := strconv.Atoi(cmd[1])

		for i := 0; i < steps; i++ {
			moveHead(dir, &h, visited)
			if moveTail(&h, &t, visited) {
				counter++
			}
		}
	}
	return counter
}

func moveHead(dir string, head *ropeEnd, visited map[string]bool) {
	head.prevX = head.x
	head.prevY = head.y

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
func moveTail(head, tail *ropeEnd, visited map[string]bool) bool {
	if !checkNeighbour(*head, *tail) {
		tail.x = head.prevX
		tail.y = head.prevY
		key := fmt.Sprintf("%v, %v", tail.x, tail.y)
		if !visited[key] {
			visited[key] = true
			return true
		}
	}
	return false
}

func checkNeighbour(head ropeEnd, tail ropeEnd) bool {
	distanceX := (head.x - tail.x)
	distanceY := (head.y - tail.y)
	if distanceX > 1 || distanceY > 1 || distanceX < -1 || distanceY < -1 {
		return false
	}
	return true
}
