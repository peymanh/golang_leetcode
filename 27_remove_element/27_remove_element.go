package removeelement

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
