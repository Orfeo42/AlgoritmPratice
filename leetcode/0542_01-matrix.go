package leetcode

// https://leetcode.com/problems/01-matrix/

func updateMatrix(mat [][]int) [][]int {
	l := len(mat)
	h := len(mat[0])
	maxDis := l + h - 2
	dp := make([][]int, l)

	for i := range dp {
		dp[i] = make([]int, h)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < h; j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			top := maxDis
			left := maxDis
			if i > 0 {
				top = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}
			dp[i][j] = min(top, left) + 1

		}
	}

	for i := l - 1; i >= 0; i-- {
		for j := h - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			bottom := maxDis
			right := maxDis
			if i < l-1 {
				bottom = dp[i+1][j]
			}
			if j < h-1 {
				right = dp[i][j+1]
			}
			dp[i][j] = min(dp[i][j], min(bottom, right)+1)
		}
	}

	return dp
}
