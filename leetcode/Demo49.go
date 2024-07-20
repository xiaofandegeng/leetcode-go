package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 定义一个map装排序后的字符串和当前字符串
	m := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		// 排序后的bytes转换为字符串
		sortStr := string(bytes)
		// 放入到map里面
		m[sortStr] = append(m[sortStr], str)
	}
	// 定义返回结构体
	res := make([][]string, 0, len(m))

	// 取出map里面的存储的同义词的集合list
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
