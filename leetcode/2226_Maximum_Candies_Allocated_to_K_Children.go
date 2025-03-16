package leetcode

import "math"

//https://leetcode.com/problems/maximum-candies-allocated-to-k-children

func MaximumCandies(candies []int, k int64) int {
	vals := values(candies)
	if k == 1 {
		return vals.max
	}
	if vals.tot < k {
		return 0
	}
	if vals.tot == k {
		return 1
	}
	minC := 1
	maxC := vals.max
	maxValue := 0
	for minC <= maxC {
		amount := minC + (maxC-minC)/2
		if amount <= maxValue {
			minC = amount + 1
			continue
		}
		if canSplit(candies, k, amount) {
			maxValue = amount
			minC = amount + 1
			continue
		}
		maxC = amount - 1
	}

	return maxValue
}

func canSplit(candies []int, k int64, amount int) bool {
	var total int64 = 0
	for _, c := range candies {
		total += int64(c) / int64(amount)
		if total >= k {
			return true
		}
	}
	return false
}

type vals struct {
	max    int
	min    int
	maxAlf int
	tot    int64
}

func values(candies []int) vals {
	result := vals{
		max:    0,
		min:    math.MaxInt,
		maxAlf: 0,
		tot:    0,
	}
	for _, c := range candies {
		result.tot += int64(c)
		if result.max < c {
			result.max = c
		}
		if result.min > c {
			result.min = c
		}
		if result.maxAlf < c/2 {
			result.maxAlf = c / 2
		}
	}
	return result
}
