package leetcode

func moveZeroes(nums []int) {
	// 1. 定义两个快慢指针
	slow, fast := 0, 0
	// 2. 循环nums
	for fast < len(nums) {
		// 3. 如果nums[fast] != 0，则将nums[fast]和nums[slow]交换
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			// 4. 移动slow指针
			slow++
		}
		// 5. 移动fast指针
		fast++
	}
}
