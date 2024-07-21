package doublePointer

func moveZeroes(nums []int) {
	// 定义一个快慢指针
	slow := 0
	fast := 0

	// 当快指针小于数组长度时
	for fast < len(nums) {
		// 如果快指针当前值不为0，则交换到前面去，那么0自然就到后面了
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
