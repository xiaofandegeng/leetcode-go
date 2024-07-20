package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 定义一个map接收排序后的str 和 string的数组
	m := make(map[string][]string)

	for _, str := range strs {
		// 将str转换为byte排序后再转换为string
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		tmp := string(bytes)
		m[tmp] = append(m[tmp], str)
	}
	// 定义返回结果集合
	res := make([][]string, 0, len(m))

	for _, v := range m {
		res = append(res, v)
	}
	return res
}
