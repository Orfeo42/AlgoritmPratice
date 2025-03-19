package leetcode

//https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i

func MinOperations(nums []int) int {
	res := 0
	numLen := len(nums) - 2
	for i := 0; i < numLen; i++ {
		if i == numLen-1 {
			if nums[i] == nums[i+1] && nums[i] == nums[i+2] {
				if nums[i] == 0 {
					res++
				}
				return res
			}
			return -1
		}
		if nums[i] != 0 {
			continue
		}
		nums[i] = 1
		nums[i+1] ^= 1
		nums[i+2] ^= 1
		res++
	}
	return res
}
