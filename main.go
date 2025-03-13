package main

import "test-go/minZeroArray"

func main() {
	//dijkstra.Test()
	//breadthFirst.TestBreadthFirst()
	//dynamicProgramming.TestBackPack()
	//dynamicProgramming.TestMaxEqualsChars()
	//dynamicProgramming.TestMaxSubString()
	//backtracking.AllSudokuCombination()
	//result := slidingDoors.NumberOfSubstrings("abcabc")
	//fmt.Println(result)
	nums := []int{2, 0, 2}
	queries := [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}
	minZeroArray.MinZeroArray(nums, queries)

}
