package backtracking

import (
	"fmt"
	"strings"
)

/*
1 - Identifica il problema e le possibili scelte.
2 - Applica una scelta e verifica se è valida.
3 - Se è valida, procedi con la prossima scelta.
4 - Se arrivi a un punto morto, annulla la scelta e prova un’altra opzione.
5 - Se trovi una soluzione completa, salvala o termina il processo.

*/

func NQueensProblem(size int) {
	grid := make([]int, size)
	solve(0, grid, size)
	fmt.Println(grid)
}

func solve(row int, grid []int, size int) {
	if size == row {
		printBoard(grid, size)
		return
	}
	for col := 0; col < size; col++ {
		if !isValid(grid, row, col) {
			continue
		}
		grid[row] = col
		solve(row+1, grid, size)
		grid[row] = -1
	}
}

func isValid(grid []int, row, col int) bool {
	for prevRow := 0; prevRow < row; prevRow++ {
		prevCol := grid[prevRow]
		if prevCol == col {
			return false
		}
		dist := row - prevRow
		colTest1 := col - dist
		if prevCol == colTest1 {
			return false
		}
		colTest2 := col + dist
		if prevCol == colTest2 {
			return false
		}
	}
	return true
}

func printBoard(board []int, n int) {
	for _, col := range board {
		rowStr := strings.Repeat(" |", col) + "Q" + strings.Repeat("| ", n-col-1)
		fmt.Println("|" + rowStr + "|")
	}
	fmt.Println()
}
