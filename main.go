package main

import "fmt"

var directory = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

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
			if grid[i][j] == 1 && !vist[i][j] {
				cnt++
				bfs(grid, &vist, i, j)
			}
		}
	}
	fmt.Println(cnt)
}

func bfs(grid [][]int, vist *[][]bool, x, y int) {
	queue := make([][]int, 0) // 辅助队列
	queue = append(queue, []int{x, y})
	for len(queue) > 0 {
		e := queue[0] // 出队
		queue = queue[1:]
		for _, dir := range directory {
			newX, newY := e[0]+dir[0], e[1]+dir[1]
			if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == 1 && (*vist)[newX][newY] == false {
				queue = append(queue, []int{newX, newY})
				(*vist)[newX][newY] = true
			}
		}
	}
}
