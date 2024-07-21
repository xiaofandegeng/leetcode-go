package doublePointer

func maxArea(height []int) int {
	maxSize := 0

	fast := len(height) - 1
	slow := 0

	for slow < fast {
		// 当前高度
		curHeight := min(height[slow], height[fast])
		curS := (fast - slow) * curHeight
		maxSize = max(curS, maxSize)

		if height[slow] < height[fast] {
			slow++
		} else {
			fast--
		}

	}

	return maxSize
}
