package main

import (
	"fmt"
	removeelement "golang_leetcode/27_remove_element"
)

func main() {
	nums1 := []int{3, 2, 2, 3}
	occur := removeelement.RemoveElement(nums1, 2)
	fmt.Printf("result array: %+v, occur: %+d\n", nums1, occur)
}
