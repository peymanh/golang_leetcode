package main

import (
	"fmt"
	mergesortedarray "golang_leetcode/88_merge_sorted_array"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	mergesortedarray.Merge(nums1, m, nums2, n)
	fmt.Printf("merged array: %+v\n", nums1)
}
