package reversewordsinastring

import (
	"slices"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
