package leetcode

//https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i

func maximumTripletValue(nums []int) int64 {
	var res int64 = 0
	maxNum := 0
	maxDif := 0
	for _, num := range nums {
		tmpRes := int64(maxDif * num)
		res = maxInt64(res, tmpRes)
		maxDif = max(maxDif, maxNum-num)
		maxNum = max(maxNum, num)
	}
	return res
}
