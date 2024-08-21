package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "grayjunzi")
	userId, err := fetchUserData(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v -> %+v\n", time.Since(start), userId)
}

func fetchUserData(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	val := ctx.Value("username")
	fmt.Println("the value = ", val)

	type result struct {
		userId string
		err    error
	}
	resultch := make(chan result, 1)

	go func() {
		res, err := fetchData()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// Done
	// 1. -> context 已经超时
	// 2. -> context 已经被手动取消
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}
}

func fetchData() (string, error) {
	time.Sleep(time.Millisecond * 100)
	return "user id 1", nil
}
