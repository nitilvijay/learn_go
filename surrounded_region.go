package main

import "fmt"

func solve(board [][]byte) {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }

    rows, cols := len(board), len(board[0])

    // Step 1: Run DFS from 'O's on the outer boundaries
    for r := 0; r < rows; r++ {
        if board[r][0] == 'O' { dfs(board, r, 0) }
        if board[r][cols-1] == 'O' { dfs(board, r, cols-1) }
    }
    for c := 0; c < cols; c++ {
        if board[0][c] == 'O' { dfs(board, 0, c) }
        if board[rows-1][c] == 'O' { dfs(board, rows-1, c) }
    }

    // Step 2: Final traversal to update the board state
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if board[r][c] == 'O' {
                board[r][c] = 'X' // Surrounded region -> flip to 'X'
            } else if board[r][c] == 'E' {
                board[r][c] = 'O' // Safe region -> restore to 'O'
            }
        }
    }
}

func dfs(board [][]byte, r, c int) {
    if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] != 'O' {
        return
    }

    board[r][c] = 'E' // Mark as visited and non-surrounded

    dfs(board, r-1, c) // Up
    dfs(board, r+1, c) // Down
    dfs(board, r, c-1) // Left
    dfs(board, r, c+1) // Right
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)

	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}
}