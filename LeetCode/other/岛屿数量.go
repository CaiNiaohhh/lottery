package main

// 岛屿数量，经典的dfs

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				ans += 1
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, x, y int)  {
	high, width := len(grid), len(grid[0])
	if x < 0 || x >= high || y < 0 || y >= width {
		return
	}
	if grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(grid, x - 1, y)
	dfs(grid, x + 1, y)
	dfs(grid, x, y - 1)
	dfs(grid, x, y + 1)
}



