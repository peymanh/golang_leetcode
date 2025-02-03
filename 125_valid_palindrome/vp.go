package validpalindrome

import (
	"strings"
	"unicode"
)

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b)) {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func IsPalindrome(s string) bool {
	s_striped := strings.Replace(strings.ToLower(strip(s)), " ", "", len(s))
	reversed_s_striped := reverse(s_striped)

	return s_striped == reversed_s_striped
}
