package main

import (
	"fmt"
	"os"
)

const (
	size      = 9
	emptyCell = '.'
)

type board [size][size]rune

func main() {
	if len(os.Args) != size+1 {
		fmt.Println("Error")
		return
	}

	var b board
	for i, arg := range os.Args[1:] {
		if len(arg) != size {
			fmt.Println("Error")
			return
		}

		for j, ch := range arg {
			b[i][j] = ch
		}
	}

	if !isValidSudoku(b) {
		fmt.Println("Error")
		return
	}

	if solveSudoku(&b) {
		printBoard(b)
	} else {
		fmt.Println("Error")
	}
}

func isValidSudoku(b board) bool {
	return isRowsValid(b) && isColsValid(b) && isBlocksValid(b)
}

func isRowsValid(b board) bool {
	for i := 0; i < size; i++ {
		if !isSetValid(b[i][:]) {
			return false
		}
	}
	return true
}

func isColsValid(b board) bool {
	for j := 0; j < size; j++ {
		col := [size]rune{}
		for i := 0; i < size; i++ {
			col[i] = b[i][j]
		}
		if !isSetValid(col[:]) {
			return false
		}
	}
	return true
}

func isBlocksValid(b board) bool {
	for i := 0; i < size; i += 3 {
		for j := 0; j < size; j += 3 {
			block := [size]rune{}
			index := 0
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					block[index] = b[i+x][j+y]
					index++
				}
			}
			if !isSetValid(block[:]) {
				return false
			}
		}
	}
	return true
}

func isSetValid(set []rune) bool {
	seen := make(map[rune]bool)
	for _, ch := range set {
		if ch != emptyCell {
			if seen[ch] {
				return false
			}
			seen[ch] = true
		}
	}
	return true
}

func solveSudoku(b *board) bool {
	row, col := -1, -1

	// Find the next empty cell
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (*b)[i][j] == emptyCell {
				row, col = i, j
				break
			}
		}
		if row != -1 {
			break
		}
	}

	// No more empty cells, puzzle solved
	if row == -1 {
		return true
	}

	for num := '1'; num <= '9'; num++ {
		if isSafe(b, row, col, num) {
			(*b)[row][col] = num
			if solveSudoku(b) {
				return true
			}
			(*b)[row][col] = emptyCell // Backtrack if the current choice is not correct
		}
	}

	return false
}

func isSafe(b *board, row, col int, num rune) bool {
	// Check if the number exists in the row
	for i := 0; i < size; i++ {
		if (*b)[i][col] == num {
			return false
		}
	}

	// Check if the number exists in the column
	for j := 0; j < size; j++ {
		if (*b)[row][j] == num {
			return false
		}
	}

	// Check if the number exists in the 3x3 block
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*b)[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

func printBoard(b board) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%c ", b[i][j])
		}
		fmt.Println()
	}
}
