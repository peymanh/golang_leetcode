package main

import (
	"fmt"

	checkstringswap "golang_leetcode/1790_check_string_swap"
)

func main() {
	s1 := "caa"
	s2 := "aaz"

	r := checkstringswap.AreAlmostEqual(s1, s2)
	fmt.Println(r)
}
