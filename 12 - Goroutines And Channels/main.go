package main

import (
	"fmt"
	"time"
)

func main() {
	// 开启协程
	go fetchResource(1)

	// 调用匿名函数
	go func() {
		result := fetchResource(1)
		fmt.Println(result)
	}()
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
