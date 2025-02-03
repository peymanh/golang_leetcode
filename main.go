package main

import (
	"fmt"

	reversewordsinastring "golang_leetcode/151_reverse_words_in_a_string"
)

func main() {
	s := "a good   example"
	r := reversewordsinastring.ReverseWords(s)
	fmt.Println(r)
}
