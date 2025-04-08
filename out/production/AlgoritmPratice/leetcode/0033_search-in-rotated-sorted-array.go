package leetcode

func Search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		current := nums[mid]
		if current == target {
			return mid
		}
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
