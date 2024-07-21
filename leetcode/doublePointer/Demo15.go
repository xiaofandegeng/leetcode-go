package doublePointer

import "sort"

func ThreeSum(nums []int) [][]int {
	// 当数组的个数小于3 直接返回nil
	if len(nums) < 3 {
		return nil
	}
	var res [][]int
	// 将nums进行排序
	sort.Ints(nums)

	// 定义左右两个指针
	for i := 0; i < len(nums); i++ {
		// 定义左右两个指针
		l := i + 1
		r := len(nums) - 1
		// 去除i 和i-1相等值的情况
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			// 求三者的和
			sum := nums[i] + nums[l] + nums[r]
			// 当小于0的时候，移动l
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				// 这是三者之和为0，满足条件，放入到结果集合里面
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 移动l和r
				l++
				r--
				// 排查移动后的l和l-1的值相同，但是在l<r的前提下进行的
				if l < r && nums[l] == nums[l-1] {
					l++
				}
				// 排查移动后的r和r+1的值相同，但是在l<r的前提下进行的
				if l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}

	}

	return res
}
