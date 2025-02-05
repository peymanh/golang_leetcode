package checkstringswap

func AreAlmostEqual(s1 string, s2 string) bool {
	type Pair struct {
		First  rune
		Second rune
	}
	pairs := []*Pair{}
	s1Chars, s2Chars := []rune(s1), []rune(s2)
	for i := 0; i < len(s1Chars); i++ {
		if s1Chars[i] != s2Chars[i] {
			pairs = append(pairs, &Pair{
				First:  s1Chars[i],
				Second: s2Chars[i],
			})
		}
	}
	return (len(pairs) == 0) || (len(pairs) == 2 && pairs[0].First == pairs[1].Second && pairs[0].Second == pairs[1].First)
}
