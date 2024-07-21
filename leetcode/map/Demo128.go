package _map

func longestConsecutive(nums []int) int {
	// 定义最大长度
	maxLen := 0

	// 将nums放入到一个map里面，如果存在就设为true
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	for k := range m {
		// 当当前值不存在比它小1的值，进入循环技术
		if !m[k-1] {
			// 设一个当前值和当前长度
			curNum := k
			curLen := 1
			for m[curNum+1] {
				curNum += 1
				curLen += 1
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return maxLen

}
