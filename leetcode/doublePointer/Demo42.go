package doublePointer

func trap(height []int) int {
	// 定义返回结果
	total := 0
	// 双指针
	left := 0
	right := len(height) - 1

	// 定义左右最大高度
	leftMax := 0
	rightMax := 0

	for left < right {
		// 更新左右最大高度
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		// 当左边最大高度小于右边最大高度，移动left，并计算储水量
		if leftMax < rightMax {
			// 左边最大高度减去左边当前高度，就是左边最大高度到左边当前位置的接水量
			total += leftMax - height[left]
			// 移动左边当前值
			left += 1
		} else {
			// 右边最大高度减去右边当前高度，就是右边最大高度到右边当前位置的接水量
			total += rightMax - height[right]
			// 移动右边当前值
			right -= 1
		}
	}

	return total
}

// 导出的 Trap 函数，用于调用未导出的 trap 函数
func Trap(height []int) int {
	return trap(height)
}
