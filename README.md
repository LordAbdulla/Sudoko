# Sudoku Solver

This is a simple Sudoku solver written in Go. The program takes a Sudoku puzzle as input through the command line and prints the completed solution, if one exists. It uses a backtracking algorithm to solve the puzzle.

## How it works

- The program expects exactly 9 strings as command-line arguments.
- Each string should be 9 characters long.
- Valid characters are digits from '1' to '9' and a dot `.` to represent empty cells.
- If the input is valid and the puzzle has one possible solution, it prints the solved board.
- If the input is invalid or the puzzle is unsolvable, it prints `Error`.

## Usage

To run the program, use:

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
