package slidingWindow

func findAnagrams(s string, p string) []int {
	// 定义结果数组，用于存放所有字母异位词的起始索引
	var res []int
	// 定义两个长度为26的数组，分别存放 s 和 p 中字符的频率
	sInt := make([]int, 26)
	pInt := make([]int, 26)
	// 定义左右指针
	l, r := 0, 0

	// 统计 p 中每个字符的出现次数
	for i := 0; i < len(p); i++ {
		ch := p[i]
		pInt[ch-'a']++
	}

	// 滑动窗口遍历字符串 s
	for r < len(s) {
		// 获取当前字符
		ch := s[r]
		// 增加当前字符在 sInt 中的计数
		sInt[ch-'a']++
		// 移动右指针
		r++

		// 如果当前字符在窗口中的数量超过在 p 中的数量，缩小窗口
		for sInt[ch-'a'] > pInt[ch-'a'] {
			sInt[s[l]-'a']--
			l++
		}

		// 当窗口大小等于 p 的长度时，记录结果
		if r-l == len(p) {
			res = append(res, l)
		}
	}

	return res
}
