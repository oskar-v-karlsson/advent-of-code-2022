package solver

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

type monkey struct {
	items       []int
	op          []string
	div         int
	tru         int
	fal         int
	inspections int
}

func calcMonkeyBusiness() int {
	monkeys := parseMonkeys()
	mostActive := 0
	secondActive := 0

	for i := 0; i < 20; i++ {
		doRound(monkeys)
	}

	for i := range monkeys {
		switch {
		case monkeys[i].inspections > mostActive:
			secondActive = mostActive
			mostActive = monkeys[i].inspections
		case monkeys[i].inspections > secondActive:
			secondActive = monkeys[i].inspections
		}
	}
	return mostActive * secondActive
}

func calcPanicedMonkeyBusiness() int {
	monkeys := parseMonkeys()
	mostActive := 0
	secondActive := 0
	modulo := 1

	for i := range monkeys {
		modulo *= monkeys[i].div
	}

	for i := 0; i < 10000; i++ {
		doPanicRound(modulo, monkeys)
	}

	for i := range monkeys {
		switch {
		case monkeys[i].inspections > mostActive:
			secondActive = mostActive
			mostActive = monkeys[i].inspections
		case monkeys[i].inspections > secondActive:
			secondActive = monkeys[i].inspections
		}
	}
	return mostActive * secondActive
}

func doRound(monkeys []monkey) {
	for i := range monkeys {
		for y := range monkeys[i].items {
			monkeys[i].inspections++
			calcWorry(monkeys[i].op, &monkeys[i].items[y])
			monkeys[i].items[y] /= 3
			if monkeys[i].items[y]%monkeys[i].div == 0 {
				monkeys = throwItem(y, i, monkeys[i].tru, monkeys)
			} else {
				monkeys = throwItem(y, i, monkeys[i].fal, monkeys)
			}
		}
		monkeys[i].items = nil
	}
}

func doPanicRound(modulo int, monkeys []monkey) {
	for i := range monkeys {
		for y := range monkeys[i].items {
			monkeys[i].inspections++
			calcWorry(monkeys[i].op, &monkeys[i].items[y])
			monkeys[i].items[y] %= modulo
			if monkeys[i].items[y]%monkeys[i].div == 0 {
				monkeys = throwItem(y, i, monkeys[i].tru, monkeys)
			} else {
				monkeys = throwItem(y, i, monkeys[i].fal, monkeys)
			}
		}
		monkeys[i].items = nil
	}
}

func throwItem(index, from, to int, monkeys []monkey) []monkey {
	monkeys[to].items = append(monkeys[to].items, monkeys[from].items[index])
	return monkeys
}

func calcWorry(op []string, item *int) {
	switch op[1] {
	case "+":
		if op[2] == "old" {
			*item = *item + *item
		} else {
			n, _ := strconv.Atoi(op[2])
			*item = *item + n
		}
	case "*":
		if op[2] == "old" {
			*item = *item * *item
		} else {
			n, _ := strconv.Atoi(op[2])
			*item = *item * n
		}
	}
}

func parseMonkeys() []monkey {
	var monkeys []monkey
	s := utils.GetScanner("./day11/input.txt")
	defer utils.CloseOS()

	for s.Scan() {
		if strings.Contains(s.Text(), "Monkey") {
			s.Scan()
			r := regexp.MustCompile(`[0-9]+`)
			items := utils.StrSlcToIntSlc(r.FindAllString(s.Text(), -1))

			s.Scan()
			_, after, _ := strings.Cut(s.Text(), "= ")
			ops := strings.Split(after, " ")

			s.Scan()
			div, _ := strconv.Atoi(r.FindString(s.Text()))

			s.Scan()
			tru, _ := strconv.Atoi(r.FindString(s.Text()))

			s.Scan()
			fal, _ := strconv.Atoi(r.FindString(s.Text()))
			monkeys = append(monkeys, monkey{items: items, op: ops, div: div, tru: tru, fal: fal, inspections: 0})
		}
	}
	return monkeys
}
