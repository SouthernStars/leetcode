package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var graph = make(map[int]map[int]int)

var sk int
var total int

type Item struct {
	path []int
	val  int
}

var ans = make([]Item, 0)
var path = make([]int, 0)
var value = 0
var cnt = 0

func main() {

	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	line1Split := strings.Split(strings.TrimSpace(line1), " ")
	sk, _ = strconv.Atoi(line1Split[0])
	total, _ = strconv.Atoi(line1Split[1])

	line2, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(line2))
	for i := 0; i < m; i++ {
		linei, _ := reader.ReadString('\n')
		lineiSplit := strings.Split(strings.TrimSpace(linei), " ")
		s, _ := strconv.Atoi(lineiSplit[0])
		d, _ := strconv.Atoi(lineiSplit[1])
		p, _ := strconv.Atoi(lineiSplit[2])
		if v, ok := graph[s]; ok {
			v[d] += p
		} else {
			graph[s] = make(map[int]int)
			graph[s][d] += p
		}
	}
	// 满足条件的列表
	path = append(path, sk)
	backtrace(sk)
	//排序结果
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].val != ans[j].val {
			return ans[i].val > ans[j].val
		}
		return cmp(ans[i].path, ans[j].path)
	})
	if len(ans) > 0 {
		fmt.Println(ans[0].val)
		for i := 0; i < len(ans[0].path); i++ {
			fmt.Printf("%d", ans[0].path[i])
			if i < len(ans[0].path)-1 {
				fmt.Printf(" ")
			}
		}
	}
}

func cmp(a []int, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	// 都相等
	return len(a) < len(b)
}

func backtrace(st int) {
	// 收集结果
	if cnt <= total {
		item := Item{
			path: append([]int{}, path...),
			val:  value,
		}
		ans = append(ans, item)
		if cnt == total {
			return
		}
	}
	// 横向遍历
	for k, v := range graph[st] {
		cnt++
		value += v
		path = append(path, k)
		backtrace(k)
		cnt--
		value -= v
		path = path[0 : len(path)-1]
	}
}
