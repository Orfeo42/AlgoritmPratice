package leetcode

import "math"

func MinCapability(nums []int, k int) int {
	if k == 1 {
		return min(nums)
	}
	lenToTest := len(nums) - ((k - 1) * 2)
	mins := make([]int, lenToTest)
	for i := 0; i < lenToTest; i++ {
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
