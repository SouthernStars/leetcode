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

import "fmt"

func main() {
	type Item struct {
		path []int
		val  int
	}

	pa := []int{}
	pa = append(pa, 1)
	item := Item{
		path: pa,
		val:  1,
	}
	fmt.Println(pa)
	fmt.Println(item)

	pa = append(pa, 2)
	fmt.Println(pa)
	fmt.Println(item)
}
