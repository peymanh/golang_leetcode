package twosum

import "slices"

func TwoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	slices.Sort(nums)
	var result []int
	for i < j {
		if nums[i]+nums[j] == target {
			result = []int{i, j}
			break
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return result
}
