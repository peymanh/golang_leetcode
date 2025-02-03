package ransomnote

func CanConstruct(ransomNote string, magazine string) bool {
	chars := make(map[rune]int)
	for _, char := range magazine {
		if _, found := chars[char]; found {
			chars[char]++
		} else {
			chars[char] = 1
		}
	}

	for _, char := range ransomNote {
		if cap, found := chars[char]; found && cap > 0 {
			chars[char]--
		} else {
			return false
		}
	}
	return true
}
