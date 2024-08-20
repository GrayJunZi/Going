package main

import "fmt"

func main() {
	// buffered channel
	// 缓冲长度最好是使用2的次方
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	for {
		msg, ok := <-msgch
		if !ok {
			return
		}

		fmt.Println(msg)
	}

	for msg := range msgch {
		fmt.Println("the msg is:", msg)
	}
}
