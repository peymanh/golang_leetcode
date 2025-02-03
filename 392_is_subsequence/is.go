package issubsequence

func isSubsequence(s string, t string) bool {
	queue := []rune(s)
	for _, char := range t {
		if len(queue) == 0 {
			continue
		}

		if queue[0] == char {
			queue = queue[1:]
		}
	}

	return len(queue) == 0
}
