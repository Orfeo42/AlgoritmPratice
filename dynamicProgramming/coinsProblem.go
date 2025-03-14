package dynamicProgramming

import mathUtility "test-go/utility"

func HowMayWays(m int, coins []int) int {
	memo := map[int]int{}
	memo[0] = 1
	for i := 1; i < m+1; i++ {
		memo[i] = 0
		for _, coin := range coins {
			subProblem := i - coin
			if subProblem < 0 {
				continue
			}
			memo[i] += memo[subProblem]
		}
	}
	return memo[m]
}

var results = map[int]int{
	0: 0,
}

func MinimumCoinsRecursive(m int, coins []int) int {
	if val, exists := results[m]; exists {
		return val
	}
	response := -1
	for _, coin := range coins {
		subProblem := m - coin
		if subProblem < 0 {
			continue
		}
		response = mathUtility.MinIntPositive(response, MinimumCoinsRecursive(subProblem, coins)+1) // +1 perchè aggiungo il soldo che ho sottratto
	}
	results[m] = response
	return response
}

func MinimumCoins(m int, coins []int) int {
	memo := make([]int, m+2)
	for i := 1; i <= m+1; i++ {
		if memo[i] == 0 {
			memo[i] = -1
		}
		for _, coin := range coins {
			subProblem := i - coin
			if subProblem < 0 {
				continue
			}
			subResult := mathUtility.MinIntPositive(memo[i], memo[subProblem]+1)
			if subResult < 1 {
				continue
			}
			memo[i] = subResult
		}
	}

	return memo[m]
}
