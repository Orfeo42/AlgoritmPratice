package leetcode

//https://leetcode.com/problems/sum-of-all-subset-xor-totals

func subsetXORSum(nums []int) int {
	or := 0
	for _, num := range nums {
		or |= num
	}
	return or << (len(nums) - 1)
}
