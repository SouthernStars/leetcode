package main

import "fmt"

// 上下左右
var direction = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &grid[i][j])
		}
	}
	cnt := 0
	// 按层顺序遍历
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				// 统计该节点上下左右的水域
				for _, dir := range direction {
					newX, newY := i+dir[0], j+dir[1]
					if newX >= 0 && newX < n && newY >= 0 && newY < m && grid[newX][newY] == 0 {
						cnt++
					}
				}
			}
		}
	}
	fmt.Println(cnt)
}
