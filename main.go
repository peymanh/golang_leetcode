package main

import (
	"fmt"
	buyandsell "golang_leetcode/121_buy_and_sell"
)

func main() {
	nums1 := []int{7, 1, 5, 3, 6, 4}
	profit := buyandsell.MaxProfit(nums1)
	fmt.Println(profit)
}
