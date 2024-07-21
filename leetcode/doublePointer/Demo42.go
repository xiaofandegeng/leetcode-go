package doublePointer

func trap(height []int) int {
	res := 0
	// 定义左右指针
	l := 0
	r := len(height) - 1

	// 移动l和r时 找到的左右的最大高度
	lMax := 0
	rMax := 0

	for l < r {
		// 获取到最新的lMax和rMax
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		// 当左边当前高度小于右边当前高度时
		if height[l] < height[r] {
			// 此时左边新增装水量为 左边最大高度减去当前高度 再移动左边值
			res += lMax - height[l]
			l++
		} else {
			// 此时右边新增装水量为 右边最大高度减去当前高度 再移动右边值
			res += rMax - height[r]
			r--
		}
	}

	return res
}

// 导出的 Trap 函数，用于调用未导出的 trap 函数
func Trap(height []int) int {
	return trap(height)
}
