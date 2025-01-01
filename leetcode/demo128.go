package leetcode

func longestConsecutive(nums []int) int {
	// 1. 定义返回结果
	maxLen := 0
	// 2. 定义一个map，key为int，value为bool
	m := make(map[int]bool)

	// 3. 循环nums，将nums中的值存入到map中
	for _, v := range nums {
		m[v] = true
	}
	// 4. 循环m，判断m中是否存在v-1，如果存在，则跳过
	for v := range m {
		if m[v-1] {
			continue
		}
		// 5. 如果不存在，则定义一个长度
		tmpLen := 1
		// 6. 循环m，判断m中是否存在v+1，如果存在，则长度+1
		for m[v+1] {
			tmpLen++
			v++
		}
		// 7. 判断长度是否大于maxLen，如果大于，则将长度赋值给maxLen
		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}

	return maxLen
}
