package dynamicProgramming

func HowManyPaths(w, h int) int {
	if w == 1 || h == 1 {
		return 1
	}

	memo := make([][]int, w)
	for i := range memo {
		memo[i] = make([]int, h)
	}

	for i := 0; i < w; i++ {
		memo[i][0] = 1
	}
	for x := 0; x < h; x++ {
		memo[0][x] = 1
	}
	for i := 1; i < w; i++ {
		for x := 1; x < h; x++ {
			memo[i][x] = memo[i-1][x] + memo[i][x-1]
		}
	}
	return memo[w-1][h-1]

}
