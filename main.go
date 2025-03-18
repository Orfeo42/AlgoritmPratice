package main

import "fmt"

func main() {
	nums := []int{135745088, 609245787, 16, 2048, 2097152}
	fmt.Println(longestNiceSubarray(nums))
	for end := 1; end < len(nums); end++ {
		fmt.Printf("start: %d, end: %d, res: %d\n", nums[end-1], nums[end], nums[end]&nums[end-1])
	}
}

func longestNiceSubarray(nums []int) int {
	bitmask := 0
	start := 0
	maxLen := 0
	for end := 0; end < len(nums); end++ {
		for bitmask&nums[end] != 0 {
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
