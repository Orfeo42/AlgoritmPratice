package main

import "fmt"

func main() {}
func countDays(days int, meetings [][]int) int {
	groups := [][]int{{meetings[0][0], meetings[0][1]}}

	for i := 1; i < len(meetings); i++ {
		val := meetings[i]
		added := false
		for _, group := range groups {
			if val[0] < group[1] {
				if val[0] < group[0] {
					group[0] = val[0]
				}
				if val[1] > group[1] {
					group[1] = val[1]
				}
				added = true
				break
			}
		}
		if !added {
			groups = append(groups, []int{val[0], val[1]})
		}

	}
	fmt.Println(groups)
	return 0
}
