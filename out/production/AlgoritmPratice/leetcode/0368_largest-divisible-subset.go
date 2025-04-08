package leetcode

import "sort"

//https://leetcode.com/problems/largest-divisible-subset

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	l := len(nums)
	size := make([]int, l)
	for i := range size {
		size[i] = 1
	}
	prev := make([]int, l)
	for i := range prev {
		prev[i] = -1
	}
	maxI := 0

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && size[i] < size[j]+1 {
				size[i] = size[j] + 1
				prev[i] = j
			}
		}
		if size[i] > size[maxI] {
			maxI = i
		}
	}
	var res []int
	for i := maxI; i > -1; i = prev[i] {
		res = append(res, nums[i])
		if prev[i] == -1 {
			break
		}
	}
	return res

}
