package main

import "fmt"

func main() {
	resultch := make(chan string)

	go func() {
		// 将数据从 channel 中读取出来
		result := <-resultch
		fmt.Println(result)
	}()

	// 将数据写入 channel 中
	resultch <- "foo"
}
