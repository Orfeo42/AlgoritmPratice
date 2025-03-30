package leetcode

//https://leetcode.com/problems/partition-labels

func partitionLabels(s string) []int {

	letterArray := [26]int{}
	for i := range letterArray {
		letterArray[i] = -1
	}

	for i, ch := range s {
		pos := int(ch - 'a')
		letterArray[pos] = i
	}

	var result []int
	start := 0
	end := -1
	for i, ch := range s {
		pos := int(ch - 'a')
		if letterArray[pos] > end {
			end = letterArray[pos]
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
