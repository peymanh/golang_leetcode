package longestcommonprefix

import (
	"fmt"
	"slices"
)

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	lengths := []int{}

	for _, str := range strs {
		lengths = append(lengths, len(str))
	}
	minLength := slices.Min(lengths)

	i := 0
	for i < minLength {
		setstr := make(map[byte]int)
		for _, str := range strs {
			if _, exists := setstr[str[i]]; exists {
				setstr[str[i]]++
			} else {
				setstr[str[i]] = 1
			}
		}

		if len(setstr) == 1 {
			fmt.Println("len found")

			for key := range setstr {
				prefix += string(key)
			}
		} else {
			break
		}
		i++
	}

	return prefix
}
