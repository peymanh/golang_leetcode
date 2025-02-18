package main

import (
	"fmt"

	rotatearray "golang_leetcode/189_rotate_array"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	r := rotatearray.Rotate(nums, k)
	fmt.Println(r)
}
