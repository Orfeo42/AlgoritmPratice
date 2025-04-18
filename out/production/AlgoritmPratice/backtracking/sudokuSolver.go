package backtracking

import "fmt"

func SolveSudoku() {
}

func AllSudokuCombination() {
	grid := [9][9]int{}
	findCombinations(grid, 0, 0)
}

func findCombinations(grid [9][9]int, row, col int) {
	if row == 9 {
		printSudoku(grid)
		return
	}

	for num := 1; num < 10; num++ {
		if !isValidRow(grid, num, row, col) {
			continue
		}
		if !isValidCol(grid, num, row, col) {
			continue
		}
		if !isValidSquare(grid, num, row, col) {
			continue
		}
		grid[row][col] = num
		if col == 8 {
			findCombinations(grid, row+1, 0)
		} else {
			findCombinations(grid, row, col+1)
		}
		grid[row][col] = -1
	}
}

func isValidRow(grid [9][9]int, num, row, col int) bool {
	for i := 0; i < col; i++ {
		numPrev := grid[row][i]
		if numPrev == num {
			return false
		}
	}
	return true
}

func isValidCol(grid [9][9]int, num, row, col int) bool {
	for i := 0; i < row; i++ {
		numPrev := grid[i][col]
		if numPrev == num {
			return false
		}
	}
	return true
}

func isValidSquare(grid [9][9]int, num, row, col int) bool {
	squareRow, squareCol := getSquare(row, col)
	squareRow = squareRow * 3
	squareCol = squareCol * 3
	for rowI := squareRow; rowI < row; rowI++ {
		for colI := squareCol; colI < squareCol+3; colI++ {
			numPrev := grid[rowI][colI]
			if numPrev == num {
				return false
			}
		}
	}
	return true
}

func getSquare(row, col int) (int, int) {
	squareRow := row / 3
	squareCol := col / 3
	return squareRow, squareCol
}

func printSudoku(grid [9][9]int) {
	for i := range grid {
		if i != 0 && i%3 == 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

/*
func printSudoku2(grid [9][9]int) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}
*/
