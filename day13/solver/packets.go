package solver

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/Zairian/advent-of-code-2022/utils"
)

func getCorrectOrder() (int, int) {
	counter1 := 0
	counter2 := 1
	s := utils.GetScanner("./day13/input.txt")
	defer utils.CloseOS()

	packets := []any{}

	for i := 0; s.Scan(); i++ {
		var first, second any
		json.Unmarshal([]byte(s.Text()), &first)
		s.Scan()
		json.Unmarshal([]byte(s.Text()), &second)
		packets = append(packets, first, second)

		if compareList(first, second) <= 0 {
			counter1 += i + 1
		}

		s.Scan()
	}

	packets = append(packets, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packets, func(i, j int) bool { return compareList(packets[i], packets[j]) < 0 })

	for i, packet := range packets {
		if fmt.Sprint(packet) == "[[2]]" || fmt.Sprint(packet) == "[[6]]" {
			counter2 *= i + 1
		}
	}

	return counter1, counter2
}

func compareList(first, second any) int {
	var firstSlice, secondSlice []any
	firstFloat, secondFloat := false, false

	switch first.(type) {
	case float64:
		firstSlice, firstFloat = []any{first}, true
	case []any, []float64:
		firstSlice = first.([]any)
	}

	switch second.(type) {
	case float64:
		secondSlice, secondFloat = []any{second}, true
	case []any, []float64:
		secondSlice = second.([]any)
	}

	if firstFloat && secondFloat {
		return int(firstSlice[0].(float64) - secondSlice[0].(float64))
	}
	for i := 0; i < len(firstSlice) && i < len(secondSlice); i++ {
		if c := compareList(firstSlice[i], secondSlice[i]); c != 0 {
			return c
		}
	}
	return len(firstSlice) - len(secondSlice)
}
