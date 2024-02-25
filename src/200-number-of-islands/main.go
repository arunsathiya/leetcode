package main

func dfs(i, j int, grid [][]byte) {
	if i > len(grid)-1 || i < 0 || j < 0 || j > len(grid[0])-1 || string(grid[i][j]) != "1" {
		return
	}
	grid[i][j] = '#'
	dfs(i+1, j, grid)
	dfs(i-1, j, grid)
	dfs(i, j+1, grid)
	dfs(i, j-1, grid)
}

func numIslands(grid [][]byte) int {
	c := 0
	for i, row := range grid {
		for j, v := range row {
			if string(v) == "1" {
				c++
				dfs(i, j, grid)
			}
		}
	}
	return c
}
