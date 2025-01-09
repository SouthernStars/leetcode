package main

import "fmt"

// 4 5
// 1 1 0 0 0
// 1 1 0 0 0
// 0 0 1 0 0
// 0 0 0 1 1
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
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !vist[i][j] {
				cnt := 0
				ise := false
				bfs(grid, &vist, &cnt, &ise, i, j)
				if !ise {
					ans += cnt
				}
			}
		}
	}
	fmt.Println(ans)
}

// 上下左右
var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func bfs(grid [][]int, vist *[][]bool, cnt *int, ise *bool, x, y int) {
	var queue [][2]int
	queue = append(queue, [2]int{x, y})
	(*vist)[x][y] = true
	*cnt++
	if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[0])-1 {
		*ise = true
	}

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:] // 出队
		for _, dir := range directions {
			newX, newY := e[0]+dir[0], e[1]+dir[1]
			if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == 1 && (*vist)[newX][newY] == false {
				(*vist)[newX][newY] = true
				*cnt++
				if newX == 0 || newY == 0 || newX == len(grid)-1 || newY == len(grid[0])-1 {
					*ise = true
				}
				queue = append(queue, [2]int{newX, newY})
			}
		}
	}
}
