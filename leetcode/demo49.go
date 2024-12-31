package leetcode

import "sort"

func GroupAnagrams(strs []string) [][]string {
	// 1. 定义一个map存放key为string，value为[]string的
	m := make(map[string][]string)
	// 2. 循环strs，转换为byte数组，然后转换string，存入map
	for _, str := range strs {
		bytes := []byte(str)
		// 3. 将bytes排序
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		// 4. 将byte数组转换为string
		sortedStr := string(bytes)
		// 5. 将排序后的string存入map里面
		m[sortedStr] = append(m[sortedStr], str)
	}
	// 6. 定义一个返回结果集
	result := make([][]string, 0, len(m))
	// 7. 循环m添加到结果集
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
