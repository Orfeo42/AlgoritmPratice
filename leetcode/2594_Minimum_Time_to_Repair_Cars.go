package leetcode

import "math"

//https://leetcode.com/problems/minimum-time-to-repair-cars

func RepairCars(ranks []int, cars int) int64 {
	var left int64 = 1
	right := int64(ranks[0] * cars * cars)
	cars64 := int64(cars)
	for left <= right {
		mid := (right + left) / 2
		if canRepair(mid, ranks, cars64) {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return left
}

func canRepair(time int64, ranks []int, cars int64) bool {
	for _, rank := range ranks {
		cars -= int64(math.Sqrt(float64(time / int64(rank))))
		if cars <= 0 {
			return true
		}
	}
	return false
}
