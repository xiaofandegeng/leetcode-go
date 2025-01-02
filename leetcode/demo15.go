package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	// 1. 定义返回结果
	res := make([][]int, 0)

	// 2. 将nums从小到大排序
	sort.Ints(nums)

	// 3. 循环nums，定义一个默认值，然后使用双指针
	for i := 0; i < len(nums)-2; i++ {
		// 4. 如果i值和i-1值相等，则跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 5. 定义左右指针
		l, r := i+1, len(nums)-1
		// 6. 循环左右指针
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			// 7. 如果sum等于0，则将结果存入到res中
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 8. 移动左右指针
				l++
				r--
				// 9. 如果sum小于0，则移动左指针
				for l < r && nums[l] == nums[l-1] && l-1 > 0 {
					l++
				}
				// 10. 如果sum大于0，则移动右指针
				for l < r && nums[r] == nums[r+1] && r+1 < len(nums) {
					r--
				}
			} else if sum < 0 {
				// 11. 如果sum小于0，则移动左指针
				l++
			} else {
				// 12. 如果sum大于0，则移动右指针
				r--
			}
		}
	}

	return res
}
