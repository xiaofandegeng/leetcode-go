package slidingWindow

func findAnagrams(s string, p string) []int {
	var res []int

	sInt := make([]int, 128)
	pInt := make([]int, 128)

	for i := 0; i < len(p); i++ {
		pInt[p[i]]++
	}

	l, r := 0, 0
	for r < len(s) {
		ch := s[r]
		sInt[ch]++
		r++

		for l < r && sInt[ch] > pInt[ch] {
			sInt[s[l]]--
			l++
		}
		if r-l == len(p) {
			res = append(res, l)
		}
	}

	return res
}
