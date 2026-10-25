package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	fmt.Println(closedIsland(grid))
}

func closedIsland(grid [][]int) int {
    row := len(grid)
    col := len(grid[0])

    count := 0

    //remove the land connected to the boundary

    for i:=0; i<row; i++{
		if grid[i][0] == 0 {
			dfs(grid, i, 0)
		}
		if grid[i][col-1] == 0 {
			dfs(grid, i, col-1)
		}
    }

    for j:=0; j<col; j++{
		if grid[0][j] == 0 {
			dfs(grid, 0, j)
		}
		if grid[row-1][j] == 0 {
			dfs(grid, row-1, j)
		}
    }
    //now compute the land from the revised 
    for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            if grid[i][j]==0 {
                count++
                dfs(grid, i, j)
            }
        }
    }
    return count
}

func dfs(grid [][]int, i int, j int) {
    if (i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 1) {
        return
    }
        
    grid[i][j] = 1

    // Visit four directions
    dfs(grid, i - 1, j); // Up
    dfs(grid, i + 1, j); // Down
    dfs(grid, i, j - 1); // Left
    dfs(grid, i, j + 1); // Right
}
