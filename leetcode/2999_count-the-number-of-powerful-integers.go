package leetcode

import (
	"strconv"
)

//https://leetcode.com/problems/count-the-number-of-powerful-integers

func numberOfPowerfulInt(start int64, finish int64, limit int, suffix string) int64 {
	countValid := func(num int64, limit int, suffix string) int64 {
		suffixNum, err := strconv.ParseInt(suffix, 10, 64)
		if err != nil {
			return 0
		}
		if num < suffixNum {
			return 0
		}
		numStr := strconv.FormatInt(num, 10)
		dp := make([][]*int64, len(numStr))
		for i := range dp {
			dp[i] = make([]*int64, 2)
		}
		var dfs func(idx int, tight bool) int64
		dfs = func(idx int, tight bool) int64 {
			if idx == len(numStr) {
				return 1
			}
			tightVal := 0
			if tight {
				tightVal = 1
			}
			if dp[idx][tightVal] != nil {
				return *dp[idx][tightVal]
			}

			var res int64 = 0
			maxDigit := 9
			if tight {
				maxDigit = int(numStr[idx] - '0')
			}

			suffixStart := len(numStr) - len(suffix)
			if idx >= suffixStart {
				suffixIdx := idx - suffixStart
				digit := int(suffix[suffixIdx] - '0')
				if digit <= maxDigit && digit <= limit {
					res += dfs(idx+1, tight && digit == maxDigit)
				}
			} else {
				for d := 0; d <= min(maxDigit, limit); d++ {
					res += dfs(idx+1, tight && d == maxDigit)
				}
			}
			dp[idx][tightVal] = &res
			return res
		}
		return dfs(0, true)
	}

	countToFinish := countValid(finish, limit, suffix)
	countToStart := countValid(start-1, limit, suffix)
	return countToFinish - countToStart
}
