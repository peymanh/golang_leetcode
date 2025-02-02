package lengthoflastword

import "strings"

func LengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
