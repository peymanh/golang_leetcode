package twosumsorted

func twoSum(numbers []int, target int) []int {

	i, j := 0, len(numbers)-1

	for i < j {
		n := numbers[i] + numbers[j]
		if n < target {
			i++
		} else if n > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}

}
