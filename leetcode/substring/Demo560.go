package substring

func subarraySum(nums []int, k int) int {
	// 前缀和
	count := 0
	preSum := 0
	m := make(map[int]int)
	m[0] = 1

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if _, ok := m[preSum-k]; ok {
			count += m[preSum-k]
		}
		m[preSum] += 1
	}

	return count
}
