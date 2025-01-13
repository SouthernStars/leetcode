package main

import "fmt"

var father []int

func main() {
	var N int
	fmt.Scanf("%d", &N)
	father = make([]int, N+1)
	Init(N)
	for i := 0; i < N; i++ {
		var start, end int
		fmt.Scanf("%d %d", &start, &end)
		if IsSameRoot(start, end) {
			fmt.Printf("%d %d\n", start, end)
			return
		} else {
			Join(start, end)
		}
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
