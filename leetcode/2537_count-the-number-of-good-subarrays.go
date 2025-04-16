package leetcode

//https://leetcode.com/problems/count-the-number-of-good-subarrays

func countGood(nums []int, k int) int64 {
	if len(nums) < 2 {
		return 0
	}

	countMap := make(map[int]int)
	var goodSubArrays int64 = 0
	var currentPairs int64 = 0
	left := 0
	k64 := int64(k)

	for right, num := range nums {
		count := countMap[num]
		currentPairs += int64(count)
		countMap[num] = count + 1

		for currentPairs >= k64 {
			goodSubArrays += int64(len(nums) - right)

			numLeft := nums[left]
			countLeft := countMap[numLeft]

			currentPairs -= int64(countLeft - 1)

			if countLeft == 1 {
				delete(countMap, numLeft)
			} else {
				countMap[numLeft] = countLeft - 1
			}
			left++
		}
	}

	return goodSubArrays
}
