package solver

import (
	"strconv"

	"github.com/Zairian/advent-of-code-2022/utils"
)

type coord struct {
	x int
	y int
}

func calcShortPath() (int, []string) {
	heightMap, startCoord, endCoord := parseInput()
	graph := utils.NewGraph()

	for i := 1; i < len(heightMap)-1; i++ {
		for y := 1; y < len(heightMap[i])-1; y++ {
			if heightMap[i-1][y] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i-1)+","+strconv.Itoa(y), 1)
			}
			if heightMap[i+1][y] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i+1)+","+strconv.Itoa(y), 1)
			}
			if heightMap[i][y-1] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i)+","+strconv.Itoa(y-1), 1)
			}
			if heightMap[i][y+1] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i)+","+strconv.Itoa(y+1), 1)
			}
		}
	}

	return graph.GetPath(strconv.Itoa(startCoord.x)+","+strconv.Itoa(startCoord.y), strconv.Itoa(endCoord.x)+","+strconv.Itoa(endCoord.y))
}

func calcMultiStartShortPath() (int, []string) {
	heightMap, startCoord, endCoord := parseInputMultiStart()
	graph := utils.NewGraph()

	for i := 1; i < len(heightMap)-1; i++ {
		for y := 1; y < len(heightMap[i])-1; y++ {
			if heightMap[i-1][y] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i-1)+","+strconv.Itoa(y), 1)
			}
			if heightMap[i+1][y] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i+1)+","+strconv.Itoa(y), 1)
			}
			if heightMap[i][y-1] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i)+","+strconv.Itoa(y-1), 1)
			}
			if heightMap[i][y+1] <= heightMap[i][y]+1 {
				graph.AddEdge(strconv.Itoa(i)+","+strconv.Itoa(y), strconv.Itoa(i)+","+strconv.Itoa(y+1), 1)
			}
		}
	}

	shortest := 999999999999
	var shortestRoute []string

	for i := range startCoord {
		temp, tempRoute := graph.GetPath(strconv.Itoa(startCoord[i].x)+","+strconv.Itoa(startCoord[i].y), strconv.Itoa(endCoord.x)+","+strconv.Itoa(endCoord.y))
		if temp < shortest && temp != 0 {
			shortest = temp
			shortestRoute = tempRoute
		}
	}

	return shortest, shortestRoute
}

func parseInput() ([][]int, coord, coord) {
	var heightMap [][]int
	var startCoord coord
	var endCoord coord

	s := utils.GetScanner("./day12/input.txt")
	defer utils.CloseOS()

	for i := 0; s.Scan(); i++ {
		heightMap = append(heightMap, []int{})
		for y, char := range s.Text() {
			switch char {
			case 'S':
				startCoord = coord{x: i, y: y}
				heightMap[i] = append(heightMap[i], 1)
			case 'E':
				endCoord = coord{x: i, y: y}
				heightMap[i] = append(heightMap[i], 26)
			default:
				heightMap[i] = append(heightMap[i], int(char)-96)
			}
		}
	}

	var edgeRow []int

	for i := range heightMap {
		heightMap[i] = append([]int{99}, heightMap[i]...)
		heightMap[i] = append(heightMap[i], 99)
	}

	for range heightMap[0] {
		edgeRow = append(edgeRow, 99)
	}

	heightMap = append([][]int{edgeRow}, heightMap...)
	heightMap = append(heightMap, edgeRow)

	startCoord.x = startCoord.x + 1
	startCoord.y = startCoord.y + 1
	endCoord.x = endCoord.x + 1
	endCoord.y = endCoord.y + 1

	return heightMap, startCoord, endCoord
}

func parseInputMultiStart() ([][]int, []coord, coord) {
	var heightMap [][]int
	var startCoord []coord
	var endCoord coord

	s := utils.GetScanner("./day12/input.txt")
	defer utils.CloseOS()

	for i := 0; s.Scan(); i++ {
		heightMap = append(heightMap, []int{})
		for y, char := range s.Text() {
			switch char {
			case 'S', 'a':
				startCoord = append(startCoord, coord{x: i, y: y})
				heightMap[i] = append(heightMap[i], 1)
			case 'E':
				endCoord = coord{x: i, y: y}
				heightMap[i] = append(heightMap[i], 26)
			default:
				heightMap[i] = append(heightMap[i], int(char)-96)
			}
		}
	}

	var edgeRow []int

	for i := range heightMap {
		heightMap[i] = append([]int{99}, heightMap[i]...)
		heightMap[i] = append(heightMap[i], 99)
	}

	for range heightMap[0] {
		edgeRow = append(edgeRow, 99)
	}

	heightMap = append([][]int{edgeRow}, heightMap...)
	heightMap = append(heightMap, edgeRow)

	for i := range startCoord {
		startCoord[i].x = startCoord[i].x + 1
		startCoord[i].y = startCoord[i].y + 1
	}

	endCoord.x = endCoord.x + 1
	endCoord.y = endCoord.y + 1

	return heightMap, startCoord, endCoord
}
