package leetcode

//https://leetcode.com/problems/minimum-index-of-a-valid-split

func MinimumIndex(nums []int) int {
	numSize := len(nums)
	dominant := 0
	dominantCount := 0
	for _, num := range nums {
		if dominantCount == 0 {
			dominant = num
			dominantCount++
			continue
		}
		if num == dominant {
			dominantCount++
			continue
		}
		dominantCount--
	}

	dominantCount = 0
	for _, num := range nums {
		if num != dominant {
			continue
		}
		dominantCount++
	}

	subCount := 0
	for i, num := range nums {
		if num == dominant {
			subCount++
		}
		if subCount <= (i+1)/2 {
			continue
		}
		if dominantCount-subCount > (numSize-i-1)/2 {
			return i
		}
	}
	return -1
}
