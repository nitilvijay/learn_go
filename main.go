package main

import "fmt"
//compiling main.go and numIslands.go together, to access the numIslands function from main.go

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid))
}