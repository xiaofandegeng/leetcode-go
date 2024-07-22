package slidingWindow

func lengthOfLongestSubstring(s string) int {
	res := 0
	tmp := make([]bool, 128)

	l, r := 0, 0

	for r < len(s) {
		// 当存在是的时候，移动l
		for tmp[s[r]] {
			tmp[s[l]] = false
			l++
		}
		tmp[s[r]] = true
		r++
		res = max(res, r-l)
	}

	return res
}
