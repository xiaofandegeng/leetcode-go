package leetcode

func longestConsecutive(nums []int) int {
	// 定义返回结构体
	maxLen := 0

	// 1. 定义一个map,存放这个值是否存在
	m := make(map[int]bool)
	// 2. 循环nums，将值存入map里面，存在这个值则为true
	for _, v := range nums {
		m[v] = true
	}
	// 3.循环map，如果存在这个值的前一个值，则跳过
	for k := range m {
		if m[k-1] {
			continue
		}
		// 4. 定义一个临时变量，用来存放当前的长度
		tmpLen := 1
		// 5. 循环map，如果存在这个值的后一个值，则临时变量+1
		for m[k+1] {
			tmpLen++
			k++
		}
		// 6. 如果临时变量大于最大长度，则将临时变量赋值给最大长度
		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}
	// 7. 返回最大长度
	return maxLen
}
