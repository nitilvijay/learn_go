package main

// import "fmt"

// func main() {
// 	grid := [][]byte{
// 		{'1', '1', '0', '0', '0'},
// 		{'1', '1', '0', '0', '0'},
// 		{'0', '0', '1', '0', '0'},
// 		{'0', '0', '0', '1', '1'},
// 	}

// 	fmt.Println(numIslands(grid))
// }

func numIslands(grid [][]byte) int {
    row := len(grid)
    col := len(grid[0])

    count := 0

    for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            if grid[i][j]=='1' {
                count++
                dfs(grid, i, j)
            }
        }
    }
    return count
}

func dfs(grid [][]byte, i int, j int) {
    if (i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0') {
                return
            }
    grid[i][j] = '0'

    // Visit four directions
    dfs(grid, i - 1, j); // Up
    dfs(grid, i + 1, j); // Down
    dfs(grid, i, j - 1); // Left
    dfs(grid, i, j + 1); // Right
}