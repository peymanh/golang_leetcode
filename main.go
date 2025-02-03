package main

import (
	"fmt"

	zigzagconversion "golang_leetcode/6_zigzag_conversion"
)

func main() {
	s := "AB"
	numRows := 1
	r := zigzagconversion.Convert(s, numRows)
	fmt.Println(r)
}
