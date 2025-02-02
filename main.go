package main

import (
	"fmt"
	removeduplicatesortedarray "golang_leetcode/26_remove_duplicate_sorted_array"
)

func main() {
	nums1 := []int{3, 2, 2, 3}
	n := removeduplicatesortedarray.RemoveDuplicates(nums1)
	fmt.Printf("result array: %v, occur: %+d\n", nums1, n)
}
