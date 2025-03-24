package leetcode

import "sort"

https: //leetcode.com/problems/count-days-without-meetings/

func CountDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	mergedMeetings := [][]int{{meetings[0][0], meetings[0][1]}}
	for i := 1; i < len(meetings); i++ {
		meeting := meetings[i]
		if meeting[0] > mergedMeetings[len(mergedMeetings)-1][1] {
			mergedMeetings = append(mergedMeetings, meeting)
			continue
		}
		mergedMeetings[len(mergedMeetings)-1][1] = maxInt(mergedMeetings[len(mergedMeetings)-1][1], meeting[1])
	}
	for _, group := range mergedMeetings {
		days -= group[1] - group[0] + 1
	}
	return days
}
