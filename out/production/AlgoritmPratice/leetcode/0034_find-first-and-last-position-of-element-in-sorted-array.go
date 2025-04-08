package leetcode

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array

func searchRange(nums []int, target int) []int {
	l := len(nums)
	start := 0
	end := l - 1
	i := -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			i = mid
			break
		}
		if nums[mid] < target {
			start = mid + 1
			continue
		}
		if nums[mid] > target {
			end = mid - 1
		}
	}
	if i == -1 {
		return []int{-1, -1}
	}
	end = -1
	for j := i; j < l; j++ {
		if nums[j] != target {
			end = j - 1
			break
		} else if j == l-1 {
			end = l - 1
		}
	}
	start = -1
	for j := i; j > -1; j-- {
		if nums[j] != target {
			start = j + 1
			break
		} else if j == 0 {
			start = 0
		}
	}

	return []int{start, end}
}
