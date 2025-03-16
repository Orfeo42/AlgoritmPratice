package leetcode

import "math"

func MinCapability(nums []int, k int) int {
	if k == 1 {
		return minFromArray(nums)
	}
	lenToTest := len(nums) - ((k - 1) * 2)
	minimi := make([]int, lenToTest)
	for i := 0; i < lenToTest; i++ {
		minimi[i] = math.MaxInt
		for j := range k {
			nextIndex := i + (j+1)*2
			if nextIndex >= len(nums) {
				break
			}
			if minimi[i] > nums[nextIndex] {
				minimi[i] = nums[nextIndex]
			}
		}
	}
	for i := 0; i < len(minimi); i++ {
		if minimi[i] == math.MaxInt {
			minimi[i] = nums[i]
			continue
		}

		if minimi[i] < nums[i] {
			minimi[i] = nums[i]
		}
	}
	return minFromArray(minimi)
}

func minFromArray(memo []int) int {
	result := math.MaxInt
	for _, num := range memo {
		if num < result {
			result = num
		}
	}
	return result
}
