package leetcode

// https://leetcode.com/problems/longest-nice-subarray

func LongestNiceSubarray(nums []int) int {
	bitmask := 0
	start := 0
	maxLen := 0
	for end := 0; end < len(nums); end++ {
		for (bitmask & nums[end]) != 0 {
			bitmask ^= nums[start]
			start++
		}
		bitmask |= nums[end]
		maxLen = maxInt(maxLen, end-start+1)
	}
	return maxLen
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
