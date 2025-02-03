package indexoffirstoccurance

func StrStr(haystack string, needle string) int {
	l, i, result := len(needle), 0, -1

	for i < len(haystack)-l {
		if haystack[i:i+l] == needle {
			result = i
			break
		}
		i++
	}
	return result
}
