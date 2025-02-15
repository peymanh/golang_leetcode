package main

import (
	"fmt"

	findnumberdistinctcolor "golang_leetcode/3160_find_number_distinct_color"
)

func main() {
	limit := 4
	queries := [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}

	r := findnumberdistinctcolor.QueryResults(limit, queries)
	fmt.Println(r)
}
