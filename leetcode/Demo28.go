package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	// 初始化
	j := -1
	next := make([]int, len(needle))
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		// 当不相等时
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		// 当达到长度
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}

	return -1
}

func getNext(next []int, needle string) {
	// 构建前缀表
	j := -1
	next[0] = j

	for i := 1; i < len(needle); i++ {
		// 当不相等的时候
		for j >= 0 && needle[i] != needle[j+1] {
			j = next[j]
		}
		// 当相等时
		if needle[i] == needle[j+1] {
			j++
		}
		next[i] = j
	}
}

func main() {
	s1 := "aabaabaaf"
	s2 := "aabaaf"
	res := strStr(s1, s2)
	fmt.Println("最后结果为：", res)
}
