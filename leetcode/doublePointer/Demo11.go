package doublePointer

func maxArea(height []int) int {
	// 定义左右两个指针
	left := 0
	right := len(height) - 1

	// z最大体积
	maxSize := 0

	for left < right {
		// 此时取两者的相对低的高度
		minHeight := min(height[left], height[right])
		// 此时面积为
		curSize := (right - left) * minHeight
		maxSize = max(maxSize, curSize)

		// 移动左右指针，谁低谁移动
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return maxSize
}
