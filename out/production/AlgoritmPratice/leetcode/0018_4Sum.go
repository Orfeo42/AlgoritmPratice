package leetcode

import "sort"

//https://leetcode.com/problems/4sum

func FourSum(nums []int, target int) [][]int {
	var result [][]int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for i2 := i + 1; i2 < n-2; i2++ {
			if i2 > i+1 && nums[i2] == nums[i2-1] {
				continue
			}
			outerSum := target - nums[i] - nums[i2]
			start := i2 + 1
			end := n - 1
			for end > start {
				innerSum := nums[start] + nums[end]
				if outerSum == innerSum {
					result = append(result, []int{nums[i], nums[i2], nums[start], nums[end]})
					start++
					end--

					for start < end && nums[start] == nums[start-1] {
						start++
					}
					for start < end && nums[end] == nums[end+1] {
						end--
					}
					continue
				}
				if outerSum < innerSum {
					end--
					continue
				}
				start++
			}
		}
	}

	return result
}
