package leetcode

//https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array

func countPairs(nums []int, k int) int {
	l := len(nums)
	res := 0
	for i, num := range nums {
		for j := i + 1; j < l; j++ {
			if num != nums[j] {
				continue
			}
			if (i*j)%k == 0 {
				res++
			}
		}
	}
	return res
}
