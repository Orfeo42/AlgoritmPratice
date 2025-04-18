package leetcode

import (
	"strconv"
	"strings"
)

//https://leetcode.com/problems/count-and-say

func countAndSay(n int) string {
	if n == 1 {
		return strconv.Itoa(1)
	}
	return rle(countAndSay(n - 1))
}

func rle(s string) string {
	var result strings.Builder
	count := 0
	for i, ch := range s {
		if i == 0 {
			count++
			continue
		}
		if ch == rune(s[i-1]) {
			count++
			continue
		}
		result.WriteString(strconv.Itoa(count))
		result.WriteByte(s[i-1])
		count = 1
	}
	if count > 0 {
		result.WriteString(strconv.Itoa(count))
		result.WriteByte(s[len(s)-1])
	}
	return result.String()
}
