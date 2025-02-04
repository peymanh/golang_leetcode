package main

import (
	"fmt"

	twosum "golang_leetcode/1_two_sum"
)

func main() {
	nums := []int{3, 2, 4}

	r := twosum.TwoSum(nums, 4)
	fmt.Println(r)
}
