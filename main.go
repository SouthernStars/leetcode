package main

import "fmt"

func main() {
	// 使用邻接表存储
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	graph := make([][]int, n+1) // 使用 1~n
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, 0)
	}
	// 访问标记数组
	visited := make([]bool, n+1)
	// 存储边
	for i := 1; i <= k; i++ {
		var s, e int
		fmt.Scanf("%d %d", &s, &e)
		graph[s] = append(graph[s], e)
	}
	cnt := 0
	dfs(graph, &visited, 1, &cnt)
	if cnt == n {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}

func dfs(graph [][]int, visited *[]bool, cur int, cnt *int) {
	(*visited)[cur] = true
	*cnt++
	for _, e := range graph[cur] {
		if !(*visited)[e] {
			dfs(graph, visited, e, cnt)
		}
	}
}
