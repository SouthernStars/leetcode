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
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && vist[i][j] == false {
				cnt++
				vist[i][j] = true
				dfs(grid, &vist, i, j)
			}
		}
	}
	fmt.Println(cnt)
}

var directory = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func dfs(grid [][]int, vist *[][]bool, x, y int) {
	for _, dir := range directory {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == 1 && (*vist)[newX][newY] == false {
			(*vist)[newX][newY] = true
			dfs(grid, vist, newX, newY)
		}
	}
}
