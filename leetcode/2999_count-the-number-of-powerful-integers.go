package leetcode

import (
	"math"
	"strconv"
)

// https://leetcode.com/problems/count-the-number-of-powerful-integers

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	suffix, _ := strconv.ParseInt(s, 10, 64)
	lenS := len(s)
	if suffix > finish {
		return 0
	}
	div := int64(math.Pow(10, float64(lenS)))
	ps := start / div
	pf := finish / div
	if finish%div >= suffix {
		pf++
	}
	if start%div > suffix {
		ps++
	}
	return getAvailNum(pf, int64(limit)) - getAvailNum(ps, int64(limit))
}

func getAvailNum(num int64, limit int64) int64 {
	if num == 0 {
		return 0
	}
	if limit == 9 {
		return num
	}
	digits := 0
	if num > 0 {
		digits = int(math.Log10(float64(num)))
	}
	div := int64(math.Pow(10, float64(digits)))
	res := int64(0)
	for i := digits; i >= 0; i-- {
		d := num / div
		if d > limit {
			return res + int64(math.Pow(float64(limit+1), float64(i+1)))
		}
		res += d * int64(math.Pow(float64(limit+1), float64(i)))
		num %= div
		div /= 10
	}
	return res
}
