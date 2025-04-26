package isomorphicstrings

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_map, t_map := [128]int{}, [128]int{}

	for i := 0; i < len(s); i++ {
		sCh, tCh := s[i], t[i]

		if s_map[sCh] == 0 && t_map[tCh] == 0 {
			s_map[sCh] = int(tCh)
			t_map[tCh] = int(sCh)
		} else if s_map[sCh] != int(tCh) || t_map[tCh] != int(sCh) {
			return false
		}
	}
	return true
}
