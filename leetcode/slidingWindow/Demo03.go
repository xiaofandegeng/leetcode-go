package slidingWindow

func lengthOfLongestSubstring(s string) int {
	// 定义结果变量，用于记录最长无重复子串的长度
	res := 0

	// 使用一个长度为26的数组来记录每个字符在当前窗口中的出现次数
	m := make([]int, 26)

	// 定义左右指针，用于维护滑动窗口
	l := 0
	r := 0

	// 滑动窗口遍历字符串
	for r < len(s) {
		// 获取当前的字符
		ch := s[r]

		// 增加当前字符的计数
		m[ch-'a']++

		// 移动右指针
		r++

		// 如果当前字符在窗口中出现的次数超过1，缩小窗口
		for m[ch-'a'] > 1 {
			// 减少左指针字符的计数
			m[s[l]-'a']--
			// 移动左指针
			l++
		}

		// 更新结果，记录当前无重复子串的最大长度
		res = max(res, r-l)
	}

	return res
}
