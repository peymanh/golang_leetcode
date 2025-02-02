package removeduplicatesortedarray

import "slices"

func RemoveDuplicates(nums []int) int {
	nums = slices.Compact(nums)
	return len(nums)
}
