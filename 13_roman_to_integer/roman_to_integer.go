package romantointeger

func RomanToInt(s string) int {
	romanVals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	calcVal := 0

	for i := 0; i < len(s); i++ {
		if i == 0 {
			calcVal += romanVals[rune(s[i])]
		} else if romanVals[rune(s[i])] > romanVals[rune(s[i-1])] {
			calcVal -= romanVals[rune(s[i-1])]
			calcVal += romanVals[rune(s[i])] - romanVals[rune(s[i-1])]
		} else {
			calcVal += romanVals[rune(s[i])]
		}

	}
	return calcVal
}
