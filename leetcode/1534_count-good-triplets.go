package leetcode

//https://leetcode.com/problems/count-good-triplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	l := len(arr)
	res := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if absInt(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < l; k++ {
				if absInt(arr[j]-arr[k]) > b {
					continue
				}
				if absInt(arr[i]-arr[k]) > c {
					continue
				}
				res++
			}
		}
	}
	return res
}
