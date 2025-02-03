package main

import (
	"fmt"

	validanagram "golang_leetcode/242_valid_anagram"
)

func main() {
	s := "abagram"
	t := "nagaram"
	r := validanagram.IsAnagram(s, t)
	fmt.Println(r)
}
