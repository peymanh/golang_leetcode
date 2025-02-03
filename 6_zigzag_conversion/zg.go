package zigzagconversion

import (
	"fmt"
	"strings"
)

func Convert(s string, numRows int) string {
	rows := []string{}
	if numRows == 1 {
		return s
	}
	for range numRows {
		rows = append(rows, "")
	}

	i, j, change := 0, 0, 1

	for i < len(s) {
		fmt.Printf("i: %d\n", i)
		fmt.Printf("j: %d\n", j)

		rows[j] += string(s[i])
		fmt.Printf("rows[j]: %s\n", rows[j])
		if (change == 1 && j == numRows-1) || (change == -1 && j == 0) {
			change *= -1
		}
		j += change
		i++
	}
	return strings.Join(rows, "")
}
