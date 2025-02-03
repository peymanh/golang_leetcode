package main

import (
	"fmt"

	isomorphicstrings "golang_leetcode/205_isomorphic_strings"
)

func main() {
	s := "paper"
	t := "title"
	r := isomorphicstrings.IsIsomorphic(s, t)
	fmt.Println(r)
}
