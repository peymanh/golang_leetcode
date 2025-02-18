package rotatearray

func Rotate(nums []int, k int) []int {
	l := len(nums)

	for i := 0; i < k; i++ {
		nums = append([]int{nums[l-1]}, nums[0:l-1]...)
	}
	return nums
}
