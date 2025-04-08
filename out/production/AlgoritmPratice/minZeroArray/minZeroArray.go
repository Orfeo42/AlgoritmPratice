package minZeroArray

func MinZeroArray(nums []int, queries [][]int) int {
	lineSweep := make([]int, len(nums)+1)
	result := 0

	for i, num := range nums {
		for lineSweep[i] < num && result < len(queries) {
			left := queries[result][0]
			right := queries[result][1]
			decrement := queries[result][2]
			result++

			if right < i {
				continue
			}
			lineSweep[max(left, i)] += decrement
			lineSweep[right+1] -= decrement
		}

		if lineSweep[i] < num {
			return -1
		}
		lineSweep[i+1] += lineSweep[i]
	}
	return result
}
