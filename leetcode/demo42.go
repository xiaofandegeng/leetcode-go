package leetcode

func trap(height []int) int {
	// 1. 定义返回结果
	res := 0

	// 2. 定义左右指针
	left, right := 0, len(height)-1

	// 3. 定义左右最大值
	leftMax, rightMax := 0, 0

	// 4. 循环height
	for left < right {
		// 5. 更新左右最大值
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		// 6. 如果leftMax < rightMax，则更新res
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}

	return res
}
