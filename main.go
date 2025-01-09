package main

import "fmt"

//5 5
//1 3 1 2 4
//1 2 1 3 2
//2 4 7 2 1
//4 5 6 1 1
//1 4 1 2 1

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	grid := make([][]int, n)
	vistf := make([][]bool, n)
	vists := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &grid[i][j])
		}
		vistf[i] = make([]bool, m)
		vists[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		// 从第一边界逆流而上
		if !vistf[i][0] {
			vistf[i][0] = true
			dfs(grid, &vistf, i, 0)
		}
		// 从第二边界逆流而上
		if !vists[i][m-1] {
			vists[i][m-1] = true
			dfs(grid, &vists, i, m-1)
		}
	}
	for i := 0; i < m; i++ {
		// 从第一边界逆流而上
		if !vistf[0][i] {
			vistf[0][i] = true
			dfs(grid, &vistf, 0, i)
		}
		// 从第二边界逆流而上
		if !vists[n-1][i] {
			vists[n-1][i] = true
			dfs(grid, &vists, n-1, i)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vistf[i][j] && vists[i][j] {
				fmt.Printf("%d %d\n", i, j)
			}
		}
	}
}

// 上下左右
var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(grid [][]int, vist *[][]bool, x, y int) {
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) {
			if grid[newX][newY] >= grid[x][y] && (*vist)[newX][newY] == false {
				(*vist)[newX][newY] = true
				dfs(grid, vist, newX, newY)
			}
		}
	}
}
