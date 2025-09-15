//package main
//
//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type Point struct {
//	x    int
//	y    int
//	dist int
//	m    int
//}
//
//func main() {
//	var w, h, m, k int
//
//	reader := bufio.NewReader(os.Stdin)
//
//	line1, _ := reader.ReadString('\n')
//
//	line1_split := strings.Split(strings.TrimSpace(line1), " ")
//	w, _ = strconv.Atoi(line1_split[0])
//	h, _ = strconv.Atoi(line1_split[1])
//
//	line2, _ := reader.ReadString('\n')
//	m, _ = strconv.Atoi(strings.TrimSpace(line2))
//	line3, _ := reader.ReadString('\n')
//	k, _ = strconv.Atoi(strings.TrimSpace(line3))
//
//	wc, hc := (w-1)/2, (h-1)/2
//	var ans []Point
//
//	var tmp int
//
//	for i := 0; i < h; i++ {
//		line, err := reader.ReadString('\n')
//		if err == io.EOF {
//			break
//		}
//		sl := strings.Split(strings.TrimSpace(line), " ")
//		w = len(sl)
//		for j := 0; j < w; j++ {
//			tmp, _ = strconv.Atoi(sl[j])
//			if tmp == m {
//				dist := abs(i-hc) + abs(j-wc)
//				ans = append(ans, Point{x: j, y: i, dist: dist, m: tmp})
//			}
//		}
//	}
//	// 排序
//	sort.Slice(ans, func(i, j int) bool {
//		if ans[i].dist != ans[j].dist {
//			return ans[i].dist < ans[j].dist
//		}
//		// dist 相等
//		if ans[i].x != ans[j].x {
//			return ans[i].x < ans[j].x
//		}
//		return ans[i].y < ans[j].y
//	})
//
//	mk := min(k, len(ans))
//	for i := 0; i < mk; i++ {
//		fmt.Printf("%d %d", ans[i].x, ans[i].y)
//		if i < mk-1 {
//			fmt.Printf(" ")
//		}
//	}
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//
////
////func main() {
////	var n, m int
////	fmt.Scan(&n, &m)
////	cases := make([][]int, n)
////	for i := 0; i < n; i++ {
////		cases[i] = make([]int, m)
////		for j := 0; j < m; j++ {
////			fmt.Scan(&cases[i][j])
////		}
////	}
////	minSize := math.MaxInt
////	snap := make([]int, m)
////	choice := make([]int, 0)
////	backtrace(0, n, &cases, &choice, &snap, &minSize)
////
////	if minSize == math.MaxInt {
////		fmt.Println(-1)
////	} else {
////		fmt.Println(minSize)
////	}
////}
////
////func backtrace(start int, end int, cases *[][]int, choice *[]int, snap *[]int, minSize *int) {
////	// 收集结果
////	if isCover(snap) {
////		*minSize = min(*minSize, len(*choice))
////		return
////	}
////	if start == end { // 什么都不做
////		return
////	}
////	// 横向遍历
////	for i := start; i < end; i++ {
////		// 选择第 i 个
////		*choice = append(*choice, i)
////		// 更新 snap
////		for j := 0; j < len((*cases)[i]); j++ {
////			(*snap)[j] += (*cases)[i][j]
////		}
////		// 纵向递归
////		backtrace(i+1, end, cases, choice, snap, minSize)
////		// 回溯
////		*choice = (*choice)[0 : len(*choice)-1]
////		for j := 0; j < len((*cases)[i]); j++ {
////			(*snap)[j] -= (*cases)[i][j]
////		}
////	}
////}
////
////func isCover(snap *[]int) bool {
////	for i := 0; i < len(*snap); i++ {
////		if (*snap)[i] == 0 {
////			return false
////		}
////	}
////	return true
////}
////

package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

type Pair struct {
	a int
	b int
}

type maxHeap []int

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func main() {

	a := []int{1, 2, 3, 6, 5}
	sort.Ints(a)
	fmt.Println(a)

	var b []Pair
	b = append(b, Pair{a: 1, b: 2})
	b = append(b, Pair{a: 2, b: 3})
	b = append(b, Pair{a: 2, b: 4})
	b = append(b, Pair{a: 0, b: 1})
	sort.Slice(b, func(i, j int) bool {
		return b[i].a < b[j].a || b[i].a == b[j].a && b[i].b > b[j].b
	})
	fmt.Println(b)

	atoi, _ := strconv.Atoi("123")
	fmt.Println(atoi)

	hp := &maxHeap{}
	heap.Init(hp)

	heap.Push(hp, 1)
	heap.Push(hp, 2)
	heap.Push(hp, 4)
	heap.Push(hp, 3)

	pop := heap.Pop(hp)
	fmt.Println(pop)
	pop = heap.Pop(hp)
	fmt.Println(pop)
	pop = heap.Pop(hp)
	fmt.Println(pop)
	pop = heap.Pop(hp)
	fmt.Println(pop)

	mp := make(map[int]interface{})
	mp[1] = interface{}

}
