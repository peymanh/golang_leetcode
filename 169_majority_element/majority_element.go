package majorityelement

import (
	"fmt"
)

func MajorityElement(nums []int) int {
	var major int
	counts := make(map[int]int)
	for _, number := range nums {
		if _, exists := counts[number]; exists {
			counts[number]++
		} else {
			counts[number] = 1
		}
	}

	for key, val := range counts {
		if val > len(nums)/2.0 {
			major = key
		}
	}
	fmt.Println(counts)
	return major
}
