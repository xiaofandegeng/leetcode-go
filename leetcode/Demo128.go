package main

func longestConsecutive(nums []int) int {
	maxLen := 0
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	for k, _ := range m {

		if !m[k-1] {
			curNum := k
			curLen := 1

			for m[curNum+1] {
				curNum += 1
				curLen += 1
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}

	return maxLen
}
