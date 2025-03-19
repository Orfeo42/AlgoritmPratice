package leetcode

func DivideArray(nums []int) bool {
	for len(nums)%2 == 1 {
		return false
	}
	memo := make(map[int]int)
	for _, num := range nums {
		memo[num]++
	}
	for _, value := range memo {
		if value%2 != 0 {
			return false
		}
	}
	return true
}
