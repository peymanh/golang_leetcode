package isomorphicstrings

import (
	"slices"
)

func IsIsomorphic(s string, t string) bool {
	s_map, t_map := make(map[byte]int), make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if _, ok := s_map[s[i]]; ok {
			s_map[s[i]]++
		} else {
			s_map[s[i]] = 1
		}

		if _, ok := t_map[t[i]]; ok {
			t_map[t[i]]++
		} else {
			t_map[t[i]] = 1
		}
	}

	s_values, t_values := []int{}, []int{}
	for _, value := range s_map {
		s_values = append(s_values, value)
	}

	for _, value := range t_map {
		t_values = append(t_values, value)
	}

	slices.Sort(s_values)
	slices.Sort(t_values)

	return slices.Equal(s_values, t_values)
}
