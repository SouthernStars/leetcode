package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var startStr, endStr string
	fmt.Scanf("%s %s", &startStr, &endStr)

	strSet := map[string]bool{}
	for i := 1; i <= n; i++ {
		var tmp string
		fmt.Scanf("%s", &tmp)
		strSet[tmp] = true
	}
	// 广度优先遍历
	vistedMap := map[string]int{}
	queue := []string{startStr}
	vistedMap[startStr] = 1
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:] // 队头出队
		path := vistedMap[top]
		for i := 0; i < len(top); i++ { // 一个一个替换字符
			newWord := []byte(top)
			for j := 0; j < 26; j++ {
				newWord[i] = 'a' + byte(j)
				word := string(newWord)
				if word == endStr {
					fmt.Println(path + 1)
					return
				}
				if strSet[word] && vistedMap[word] == 0 {
					vistedMap[word] = path + 1
					queue = append(queue, word)
				}
			}
		}
	}
	fmt.Println(0)
}
