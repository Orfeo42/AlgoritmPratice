package main

import (
	"fmt"
	"test-go/slidingDoors"
)

func main() {
	//dijkstra.Test()
	//breadthFirst.TestBreadthFirst()
	//dynamicProgramming.TestBackPack()
	//dynamicProgramming.TestMaxEqualsChars()
	//dynamicProgramming.TestMaxSubString()
	//backtracking.AllSudokuCombination()
	result := slidingDoors.NumberOfSubstrings("abcabc")
	fmt.Println(result)
}
