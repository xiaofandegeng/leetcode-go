package doublePointer

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var mySlice [][]int

	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		l := i + 1
		r := n - 1
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				var tmpSlice []int
				tmpSlice = append(tmpSlice, nums[i])
				tmpSlice = append(tmpSlice, nums[l])
				tmpSlice = append(tmpSlice, nums[r])
				mySlice = append(mySlice, tmpSlice)
				l++
				r--
				for l < r && nums[l] == nums[l-1] && l-1 > 0 {
					l++
				}
				for l < r && nums[r] == nums[r+1] && r+1 < n {
					r--
				}
			}
		}
	}
	return mySlice

}
