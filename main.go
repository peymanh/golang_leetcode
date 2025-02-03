package main

import (
	"fmt"

	ransomnote "golang_leetcode/383_ransom_note"
)

func main() {
	ransomNote := "aa"
	magazine := "ab"
	r := ransomnote.CanConstruct(ransomNote, magazine)
	fmt.Println(r)
}
