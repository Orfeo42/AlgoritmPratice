package leetcode

import "sort"

//https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid

func MinOperations2(grid [][]int, x int) int {
	mod := grid[0][0] % x

	nums := make([]int, len(grid)*len(grid[0]))
	pos := 0
	for _, row := range grid {
		for _, num := range row {
			if num%x != mod {
				return -1
			}
			nums[pos] = num / x
			pos++
		}
	}

	sort.Ints(nums)

	median := nums[len(nums)/2]
	result := 0
	for _, num := range nums {
		result += absInt(median - num)
	}
	return result
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
