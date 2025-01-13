package main

import "fmt"

var father []int

func main() {
	var N int
	fmt.Scanf("%d", &N)
	father = make([]int, N+1)
	indegree := make([]int, N+1)
	edges := make([][2]int, N)

	// 统计每个结点的入度
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &edges[i][0], &edges[i][1])
		indegree[edges[i][1]]++
	}
	// 统计入度为2 的点及关联的边，按照逆序
	var twoEdges []int
	for i := N - 1; i >= 0; i-- {
		if indegree[edges[i][1]] == 2 {
			twoEdges = append(twoEdges, i)
		}
	}

	if len(twoEdges) == 0 { // 存在环
		Init(N)
		for i := 0; i < N; i++ {
			if IsSameRoot(edges[i][0], edges[i][1]) {
				fmt.Printf("%d %d\n", edges[i][0], edges[i][1])
				return
			} else {
				Join(edges[i][0], edges[i][1])
			}
		}
	} else { // 存在入度为 2 的点
		deleteEdge := twoEdges[0]
		Init(N)
		for i := 0; i < N; i++ {
			if i == deleteEdge {
				continue
			}
			if IsSameRoot(edges[i][0], edges[i][1]) { // 删除边后，构成了环，因此删错了, 修改删除边
				deleteEdge = twoEdges[1]
				break
			} else {
				Join(edges[i][0], edges[i][1])
			}
		}
		fmt.Printf("%d %d\n", edges[deleteEdge][0], edges[deleteEdge][1])
	}
}

// Init 初始化并查集，编号1~n
func Init(n int) {
	for i := 1; i <= n; i++ {
		father[i] = i
	}
}

// FindRoot 寻找 u 的根
func FindRoot(u int) int {
	if u == father[u] {
		return u
	}
	father[u] = FindRoot(father[u])
	return father[u]
}

func Join(u, v int) {
	u = FindRoot(u)
	v = FindRoot(v)
	if u != v {
		father[v] = u
	}
}

func IsSameRoot(u, v int) bool {
	return FindRoot(u) == FindRoot(v)
}
