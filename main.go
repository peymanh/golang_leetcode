package main

import (
	"fmt"
	majorityelement "golang_leetcode/169_majority_element"
)

func main() {
	nums1 := []int{3, 2, 3, 2, 2, 2}
	major := majorityelement.MajorityElement(nums1)
	fmt.Println(major)
}
