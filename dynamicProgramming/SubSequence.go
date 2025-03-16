package dynamicProgramming

import (
	mathUtility "test-go/utility"
)

func LongestCommonSubsequence(a, b string) int {
	memo := longestCommonSubMemo(a, b)
	return memo[len(a)-1][len(b)-1]

}

func LongestCommonSubToString(a, b string) string {
	memo := longestCommonSubMemo(a, b)
	return reconstructSubString(memo, a, b)
}

func reconstructSubString(memo [][]int, a, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	result := ""

	for i >= 0 && j >= 0 {
		if a[i] == b[j] {
			result = string(a[i]) + result
			i--
			j--
			continue
		}
		if memo[i-1][j] > memo[i][j-1] {
			i--
			continue
		}
		j--
	}
	return result
}

func longestCommonSubMemo(a, b string) [][]int {
	memo := make([][]int, len(a))

	for i := range memo {
		memo[i] = make([]int, len(b))
	}

	for ai := range a {
		for bi := range b {
			if a[ai] == b[bi] {
				if ai == 0 || bi == 0 {
					memo[ai][bi] = 1
					continue
				}
				memo[ai][bi] = memo[ai-1][bi-1] + 1
				continue
			}
			valA := 0
			valB := 0
			if ai > 0 {
				valA = memo[ai-1][bi]
			}
			if bi > 0 {
				valB = memo[ai][bi-1]
			}
			memo[ai][bi] = mathUtility.MaxInt(valA, valB)
		}
	}

	return memo
}
