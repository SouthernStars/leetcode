package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second) // 模拟生产数据
		fmt.Println("producer 生产数据", i)
		ch <- i
	}
	close(ch)
}

func main() {

}
