package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortStr := string(s)
		mp[sortStr] = append(mp[sortStr], str)
	}
	ans := make([][]string, 0, len(mp))

	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
