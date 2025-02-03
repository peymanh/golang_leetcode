package main

import (
	"fmt"

	indexoffirstoccurance "golang_leetcode/28_index_of_first_occurance"
)

func main() {
	haystack, needle := "sdadbutsad", "sad"
	r := indexoffirstoccurance.StrStr(haystack, needle)
	fmt.Println(r)
}
