package utils

import "strconv"

func StrSlcToIntSlc(s []string) []int {
	ints := make([]int, len(s))

	for i, s := range s {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
