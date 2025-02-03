package main

import (
	"fmt"

	longestcommonprefix "golang_leetcode/14_longest_common_prefix"
)

func main() {
	strs := []string{"cir", "car"}
	prefix := longestcommonprefix.LongestCommonPrefix(strs)
	fmt.Println(prefix)
}
