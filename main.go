package main

import "fmt"

// 使用 path 记录一条路径， res 记录全部可能的路径
var path []int
var res [][]int

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	// 编号 1~n, 申请邻接矩阵存储图 graph
	graph := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, 0)
	}
	// 处理输入，将边 信息保存到 graph 中   (i,i) 之间默认是没有边的
	for i := 0; i < m; i++ {
		var start, end int
		fmt.Scanf("%d %d", &start, &end)
		graph[start] = append(graph[start], end)
	}
	// 从编号为 1 开始遍历
	path = append(path, 1)
	dfs(graph, 1, n)

	// 输出结果
	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		for _, v := range res {
			for i := 0; i < len(v)-1; i++ {
				fmt.Printf("%d ", v[i])
			}
			fmt.Println(v[len(v)-1])
		}
	}
}

func dfs(graph [][]int, start int, n int) {
	// 结束条件，收集结果
	if start == n {
		// 将 path 存入 res 中
		res = append(res, append([]int{}, path...))
		return
	}
	for _, v := range graph[start] {
		path = append(path, v) // 将 v 加入 path, 并从 v 继续 dfs
		dfs(graph, v, n)
		// 回溯
		path = path[:len(path)-1]
	}

}
