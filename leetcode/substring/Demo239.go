package substring

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)

	lMax := make([]int, n)
	rMax := make([]int, n)

	for i := 0; i < n; i++ {
		if i%k == 0 {
			lMax[i] = nums[i]
		} else {
			lMax[i] = max(lMax[i-1], nums[i])
		}
		j := n - 1 - i
		if j%k == k-1 || j == n-1 {
			rMax[j] = nums[j]
		} else {
			rMax[j] = max(rMax[j+1], nums[j])
		}
	}

	for i := 0; i < n-k+1; i++ {
		j := i + k - 1
		maxV := max(rMax[i], lMax[j])
		ans[i] = maxV
	}

	return ans
}

func Main() {
	// 测试示例
	nums := []int{1, 3, -1, -3, -4, 3, 6, 7}
	k := 3
	fmt.Println("输入:", nums)
	fmt.Println("滑动窗口大小:", k)
	fmt.Println("滑动窗口最大值:", maxSlidingWindow(nums, k))
}
