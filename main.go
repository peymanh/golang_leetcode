package main

import (
	"fmt"

	validpalindrome "golang_leetcode/125_valid_palindrome"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	r := validpalindrome.IsPalindrome(s)
	fmt.Println(r)
}
