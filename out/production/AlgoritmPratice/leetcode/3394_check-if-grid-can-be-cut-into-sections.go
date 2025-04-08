package leetcode

import (
	"slices"
)

// https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections

type rectSize struct{ start, end int }

func checkValidCuts(n int, rectangles [][]int) bool {
	x := make([]rectSize, len(rectangles))
	y := make([]rectSize, len(rectangles))
	for i, rec := range rectangles {
		x[i] = rectSize{rec[0], rec[2]}
		y[i] = rectSize{rec[1], rec[3]}
	}

	return check(x) || check(y)
}

func check(rectangles []rectSize) bool {
	slices.SortFunc(rectangles, func(a, b rectSize) int {
		return a.start - b.start
	})

	rectCount := 0
	top := 0
	for _, ret := range rectangles {
		if ret.start >= top {
			rectCount++
		}
		if rectCount > 2 {
			return true
		}
		top = maxInt(top, ret.end)
	}
	return false
}
