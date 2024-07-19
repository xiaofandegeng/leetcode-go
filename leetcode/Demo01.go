package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		res := target - num
		if j, exist := m[res]; exist {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
