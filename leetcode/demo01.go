package leetcode

func twoSum(nums []int, target int) []int {
	// 1. 定义一个map
	m := make(map[int]int)

	// 2. 循环整个nums
	for k, v := range nums {
		// 3. 在map查找是否存在target-v的值
		res := target - v
		// 4. 如果存在，则返回索引
		if j, ok := m[res]; ok {
			return []int{j, k}
		}
		// 5. 不存在，将当前的值放入到map里面
		m[v] = k
	}
	return nil
}
