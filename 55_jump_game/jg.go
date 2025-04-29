package jumpgame

func CanJump(nums []int) bool {
	gas := 0

	for _, n := range nums {
		if gas < 0 {
			return false
		}

		if gas < n {
			gas = n
		}

		gas -= 1
	}
	return true
}
