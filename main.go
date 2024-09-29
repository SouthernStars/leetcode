package main

import "fmt"

func main() {
	s1 := " abc bsss fa "
	s2 := trim_space(s1)
	s3 := reverse(s2)
	s4 := word_reverse(s3)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func reverseWords(s string) string {
	// step1: trim_space()去除空格
	s_trim := trim_space(s)
	// step2: 句子翻转
	s_rev := reverse(s_trim)
	// 单词翻转
	word_rev := word_reverse(s_rev)
	return word_rev
}

func word_reverse(s string) string {
	bytes := ([]byte)(s)
	n := len(s)
	start, end := 0, 0
	// a baa fff wss
	for end <= n {
		if end == n || s[end] == ' ' {
			low, high := start, end-1
			for low <= high {
				bytes[low], bytes[high] = bytes[high], bytes[low]
				low++
				high--
			}
			start, end = end+1, end+1
		} else {
			end++
		}
	}
	return (string)(bytes)
}

func reverse(s string) string {
	bytes := ([]byte)(s)
	low, high := 0, len(s)-1
	for low <= high {
		bytes[low], bytes[high] = bytes[high], bytes[low]
		low++
		high--
	}
	return (string)(bytes)
}

func trim_space(s string) string {
	bytes := ([]byte)(s)
	// __a_b_c__
	idx, n := 0, len(s)

	// 去除重复空格
	for i := 0; i < n; i++ {
		if bytes[i] != ' ' {
			bytes[idx] = bytes[i]
			idx++
		} else { // ' '
			if idx == 0 {
				continue
			} else { // idx > 0
				if bytes[idx-1] != ' ' {
					bytes[idx] = bytes[i]
					idx++
				}
			}
		}
	}
	if idx == 0 {
		return ""
	}
	// 去除后导字符串
	if idx > 1 && bytes[idx-1] == ' ' {
		idx--
	}
	return (string)(bytes[0:idx])
}
