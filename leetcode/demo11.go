package leetcode

func maxArea(height []int) int {
	// 1. 定义结果
	maxSize := 0
	// 2. 定义左右指针
	left, right := 0, len(height)-1

	// 3. 循环height
	for left < right {
		// 4. 获取左右指针的最小值
		tmpHeight := min(height[left], height[right])
		// 5. 此时的面积为
		tmpSize := tmpHeight * (right - left)

		// 6. 判断面积是否大于maxSize
		if tmpSize > maxSize {
			maxSize = tmpSize
		}
		// 7.移动左右指针，移动左右指针的条件为，移动左右指针的最小值
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxSize
}
