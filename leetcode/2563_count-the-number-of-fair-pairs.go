package leetcode

import "sort"

//https://leetcode.com/problems/count-the-number-of-fair-pairs

func binarySearchLowerBound(nums []int, startIdx int, target int) int {
	start := startIdx
	end := len(nums)
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func binarySearchUpperBound(nums []int, startIdx int, target int) int {
	start := startIdx
	end := len(nums)
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	l := len(nums)
	var res int64 = 0

	for i := 0; i < l-1; i++ {
		num := nums[i]

		targetLower := lower - num
		targetUpper := upper - num

		start := binarySearchLowerBound(nums, i+1, targetLower)
		endPlus1 := binarySearchUpperBound(nums, i+1, targetUpper)
		end := endPlus1 - 1

		if start < l && end >= start {
			count := int64(end - start + 1)
			res += count
		}
	}

	return res
}
