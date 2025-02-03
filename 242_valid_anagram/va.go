package validanagram

import "slices"

func IsAnagram(s string, t string) bool {
	s_char := []rune(s)
	t_char := []rune(t)

	slices.Sort(s_char)
	slices.Sort(t_char)

	return slices.Equal(s_char, t_char)
}
