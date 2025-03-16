package leetcode

import "math"

func minCapability(nums []int, k int) int {
	if len(nums)-2 <= 1 {
		return min(nums)
	}
	mins := make([]int, len(nums)-2)
	for i := 0; i < len(nums)-2; i++ {
		mins[i] = math.MaxInt
		for j := range k {
			nextIndex := i + (j+1)*2
			if nextIndex >= len(nums) {
				break
			}
			if mins[i] > nums[nextIndex] {
				mins[i] = nums[nextIndex]
			}
		}
	}
	for i := 0; i < len(mins); i++ {
		if mins[i] == math.MaxInt {
			mins[i] = nums[i]
			continue
		}

		if mins[i] < nums[i] {
			mins[i] = nums[i]
		}
	}
	return min(mins)
}

func min(memo []int) int {
	result := math.MaxInt
	for _, num := range memo {
		if num < result {
			result = num
		}
	}
	return result
}
