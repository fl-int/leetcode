package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	/*
		Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
		- Each row must contain the digits 1-9 without repetition.
		- Each column must contain the digits 1-9 without repetition.
		- Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

		Note:
		- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
		- Only the filled cells need to be validated according to the mentioned rules.
	*/

	for m := 0; m < 3; m++ {
		for i := 0; i < 9; i++ {
			bits := int16(0)
			for j := 0; j < 9; j++ {
				ii := i
				jj := j
				if m == 1 {
					ii = j
					jj = i
				}
				if m == 2 {
					ii = (j % 3)
					jj = j + i%3
				}

				b := board[ii][jj]
				if b == '.' {
					continue
				}

				bit := int16(1 << b)
				if bits&bit == bit {
					return false
				}

				bits |= bit
			}
		}
	}

	return true
}

func isValidSudokuTest() {
	a := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(a)
	k := isValidSudoku(a)
	fmt.Println(k)
}
