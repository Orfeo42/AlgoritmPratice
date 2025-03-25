package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections

func CheckValidCuts(n int, rectangles [][]int) bool {
	fmt.Println(GroupX(rectangles))
	return false
}

func GroupX(rectangles [][]int) [][]int {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][0] < rectangles[j][0]
	})
	var group [][]int

	for _, rectangle := range rectangles {
		if len(group) == 0 {
			group = append(group, []int{rectangle[0], rectangle[2]})
			continue
		}
		if rectangle[2] <= group[len(group)-1][1] {
			//rettangolo nella stessa area di x di un altro
			continue
		}
		if rectangle[0] < group[len(group)-1][1] {
			//SE il rettangolo inizia nella stessa area x ma non alla fine prendo il più "ALTO"
			group[len(group)-1][1] = maxInt(rectangle[2], group[len(group)-1][1])
			continue
		}
		//group[len(group)-1][1] = maxInt(group[len(group)-1][1], rectangle[2])
	}
	return group
}

func GroupY(rectangles [][]int) [][]int {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][1] < rectangles[j][1]
	})
	var group [][]int

	for _, rectangle := range rectangles {
		if len(group) == 0 {
			group = append(group, []int{rectangle[1], rectangle[3]})
			continue
		}
		if rectangle[3] <= group[len(group)-1][1] {
			//rettangolo nella stessa area di x di un altro
			continue
		}
		if rectangle[1] < group[len(group)-1][1] {
			//SE il rettangolo inizia nella stessa area x ma non alla fine prendo il più "ALTO"
			group[len(group)-1][1] = maxInt(rectangle[3], group[len(group)-1][1])
			continue
		}
		//group[len(group)-1][1] = maxInt(group[len(group)-1][1], rectangle[3])
	}
	return group
}
