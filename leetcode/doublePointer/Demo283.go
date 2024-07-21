package doublePointer

func moveZeroes(nums []int) {
	// 当nums的长度小于等于1 直接返回源数组
	if len(nums) <= 1 {
		return
	}
	fast := 0
	slow := 0
	n := len(nums)

	for fast < n {
		if nums[fast] != 0 {
			tmp := nums[fast]
			nums[fast] = nums[slow]
			nums[slow] = tmp

			slow += 1
		}
		fast += 1
	}
}
