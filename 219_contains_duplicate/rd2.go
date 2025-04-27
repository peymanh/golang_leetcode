package containsduplicate

func containsNearbyDuplicate(nums []int, k int) bool {

	reverseMap := map[int]int{}

	for i, v := range nums {
		if last, found := reverseMap[v]; !found {
			reverseMap[v] = i
		} else {
			if i-last >= k {
				return true
			} else {
				reverseMap[v] = last
			}
		}
	}
	return false
}
