package substring

import "math"

func minWindow(s string, t string) string {
	l, r := 0, 0
	count := len(t)
	size := math.MaxInt32
	start := 0
	m := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		m[t[i]] += 1
	}

	for r < len(s) {
		ch := s[r]
		if m[ch] > 0 {
			count -= 1
		}
		m[ch] -= 1

		if count == 0 {
			for l < r && m[s[l]] < 0 {
				m[s[l]] += 1
				l += 1
			}
			if r-l+1 < size {
				size = r - l + 1
				start = l
			}
			m[s[l]] += 1
			count += 1
			l += 1

		}
		r += 1
	}

	if size == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+size]
	}

}
