package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isInitialBoardValid(board [9][9]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			ch := board[row][col]
			if ch != '.' {
				board[row][col] = '.'
				if !isValid(board, row, col, ch) {
					return false
				}
				board[row][col] = ch
			}
		}
	}
	return true
}

func printBoard(board [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(board[i][j])
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func isValid(board [9][9]rune, row, col int, ch rune) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == ch || board[i][col] == ch {
			return false
		}
		boxRow := 3*(row/3) + i/3
		boxCol := 3*(col/3) + i%3
		if board[boxRow][boxCol] == ch {
			return false
		}
	}
	return true
}

func solve(board *[9][9]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for ch := '1'; ch <= '9'; ch++ {
					if isValid(*board, row, col, ch) {
						board[row][col] = ch
						if solve(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func printError() {
	errorMsg := "Error\n"
	for _, r := range errorMsg {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) != 10 {
		printError()
		return
	}

	var board [9][9]rune

	for i := 0; i < 9; i++ {
		line := os.Args[i+1]
		if len(line) != 9 {
			printError()
			return
		}
		for j := 0; j < 9; j++ {
			ch := rune(line[j])
			if (ch >= '1' && ch <= '9') || ch == '.' {
				board[i][j] = ch
			} else {
				printError()
				return
			}
		}
	}

	if !isInitialBoardValid(board) {
		printError()
		return
	}

	if solve(&board) {
		printBoard(board)
	} else {
		printError()
	}
}
