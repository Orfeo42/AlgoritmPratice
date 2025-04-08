package leetcode

//https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct

func minimumOperations(nums []int) int {
	var dp [101]bool
	l := len(nums)
	for i := l - 1; i > -1; i-- {
		if !dp[nums[i]] {
			dp[nums[i]] = true
			continue
		}
		return i/3 + 1
	}
	return 0
}
