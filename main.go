package main

import (
	"fmt"
	lengthoflastword "golang_leetcode/58_length_of_last_word"
)

func main() {
	s := "asd asdasdasdasd    asd   "
	l := lengthoflastword.LengthOfLastWord(s)
	fmt.Println(l)
}
