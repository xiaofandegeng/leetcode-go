package ordinaryArray

import "math"

func maxSubArray(nums []int) int {
	count := 0
	maxCount := math.MinInt32

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		count += num
		maxCount = max(maxCount, count)

		if count < 0 {
			count = 0
		}
	}

	return maxCount
}
