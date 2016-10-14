package helpers

import (
	"strconv"
)

func Sum(s1 string, s2 string) int {
	i, _ := strconv.Atoi(s1)
	j, _ := strconv.Atoi(s2)

	return i + j
}
