package sum

import "slices"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // To prevent the repeat
		}
		left, right := i+1, len(nums)-1
		target := -nums[i]
		sum := nums[left] + nums[right]
		for left < right {
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				tuple := []int{nums[i], nums[left], nums[right]}
				res = append(res, tuple)
				left++
				right--
				for left < right && nums[left] == nums[left-1] { // in case we have duplicate numbers
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return res
}
