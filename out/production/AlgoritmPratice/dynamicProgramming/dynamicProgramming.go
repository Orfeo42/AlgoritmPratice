package dynamicProgramming

import (
	"fmt"
	"test-go/utility"
)

type Item struct {
	Size  int
	Value int
}

func TestBackPack() {
	items := []Item{
		{Size: 4, Value: 3000},
		{Size: 3, Value: 2000},
		{Size: 1, Value: 1500},
	}
	fmt.Println(backPackProblem(items, 4))
}

func TestMaxEqualsChars() {
	s1 := "Fish"
	s2 := "Fosh"
	fmt.Printf("TestMaxSubString %d\n", maxEqualsChars(s1, s2))
}

func maxEqualsChars(s1, s2 string) int {
	rows, cols := len(s1), len(s2)
	grid := makeMatrix(rows, cols)
	for i1, ch1 := range s1 {
		for i2, ch2 := range s2 {
			oldValue1 := 0
			oldValue2 := 0
			if i2 > 0 {
				oldValue2 = grid[i1][i2-1]
			}
			if i1 > 0 {
				oldValue1 = grid[i1-1][i2]
			}
			oldValue := mathUtility.MaxInt(oldValue2, oldValue1)
			if ch1 == ch2 {
				oldValue++
			}
			grid[i1][i2] = oldValue
		}
	}
	fmt.Println(grid)
	return grid[rows-1][cols-1]
}

func TestMaxSubString() {
	s1 := "Fisho"
	s2 := "Foish"
	fmt.Printf("TestMaxSubString %d\n", maxSubString(s1, s2))
}

func maxSubString(s1, s2 string) int {
	rows, cols := len(s1), len(s2)
	grid := makeMatrix(rows, cols)
	for i1, ch1 := range s1 {
		for i2, ch2 := range s2 {
			oldValue := 0
			if i2 > 0 && i1 > 0 {
				oldValue = grid[i1-1][i2-1]
			}
			if ch1 == ch2 {
				grid[i1][i2] = oldValue + 1
			}
		}
	}
	fmt.Println(grid)
	return maxMatrix(grid)
}

func backPackProblem(items []Item, maxSize int) int {
	rows, cols := len(items), maxSize
	grid := makeMatrix(rows, cols)
	for iC, item := range items {
		for i := 0; i < maxSize; i++ {
			empty := i + 1 - item.Size
			if empty < 0 {
				grid[iC][i] = 0
				continue
			}
			newValue := item.Value
			oldValue := 0
			if iC > 0 {
				newValue += grid[iC-1][empty]
				oldValue = grid[iC-1][i]
			}
			grid[iC][i] = mathUtility.MaxInt(newValue, oldValue)
		}
	}
	return grid[rows-1][cols-1]
}

func makeMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func maxMatrix(matrix [][]int) int {
	maximus := 0
	for subMatrix := range matrix {
		for i := range subMatrix {
			if i > maximus {
				maximus = i
			}
		}
	}
	return maximus
}
