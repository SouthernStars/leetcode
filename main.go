package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	grid := make([][]int, n)
	vist := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		vist[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &grid[i][j])
		}
	}
	// 统计非孤岛，并且标记
	cnt := 0
	for i := 0; i < n; i++ {
		if grid[i][0] == 1 && !vist[i][0] {
			cnt++
			vist[i][0] = true
			dfs(grid, &vist, &cnt, i, 0)
		}
		if grid[i][m-1] == 1 && !vist[i][m-1] {
			cnt++
			vist[i][m-1] = true
			dfs(grid, &vist, &cnt, i, m-1)
		}
	}
	for i := 0; i < m; i++ {
		if grid[0][i] == 1 && !vist[0][i] {
			cnt++
			vist[0][i] = true
			dfs(grid, &vist, &cnt, 0, i)
		}
		if grid[n-1][i] == 1 && !vist[n-1][i] {
			cnt++
			vist[n-1][i] = true
			dfs(grid, &vist, &cnt, n-1, i)
		}
	}
	// 然后再遍历一遍，统计孤岛
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !vist[i][j] {
				ans++
				vist[i][j] = true
				dfs(grid, &vist, &ans, i, j)
			}
		}
	}
	fmt.Println(ans)
}

// 上下左右
var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(grid [][]int, vist *[][]bool, cnt *int, x, y int) {
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == 1 && (*vist)[newX][newY] == false {
			(*vist)[newX][newY] = true
			*cnt++
			dfs(grid, vist, cnt, newX, newY)
		}
	}
}
