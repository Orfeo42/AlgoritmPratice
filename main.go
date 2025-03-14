package main

import (
	"fmt"
	"test-go/dynamicProgramming"
)

func main() {
	//result1 := dynamicProgramming.MinimumCoinsRecursive(13, []int{2, 4, 5})
	//fmt.Println(result1)
	//result2 := dynamicProgramming.MinimumCoins(13, []int{2, 4, 5})
	//fmt.Println(result2)
	result3 := dynamicProgramming.HowManyPaths(100, 100)
	fmt.Println(result3)
}
