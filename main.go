package main

import (
	"fmt"
	"test-go/dynamicProgramming"
)

func main() {
	result := dynamicProgramming.LongestCommonSubsequence("Pippo", "Pippo")
	fmt.Println(result)
	result2 := dynamicProgramming.LongestCommonSubToString("Pippo", "Pippo")
	fmt.Println(result2)
}
