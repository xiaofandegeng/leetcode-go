package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 1. 用一个key为string，value为[]string的map存储排序后的字符串
	m := make(map[string][]string)
	// 2. 循环strs，排序后存入m中
	for _, str := range strs {
		// 3. 将str转换为bytes数组
		bytes := []byte(str)

		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		// 4. 将排序后的bytes转换为string
		sortStr := string(bytes)
		// 5. 排序后的string存入map里面
		m[sortStr] = append(m[sortStr], str)
	}
	// 6. 定义返回结构体
	result := make([][]string, 0, len(m))
	// 7. 从map里面取出值放入结构体里面
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
