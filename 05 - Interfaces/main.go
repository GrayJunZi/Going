package main

import "fmt"

// 定义接口
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDBNumberStore struct {
}

func (s MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (s MongoDBNumberStore) Put(number int) error {
	fmt.Printf("store the number %d into the MongoDB database.\n", number)
	return nil
}

type PostgresNumberStore struct {
}

func (s PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (s PostgresNumberStore) Put(number int) error {
	fmt.Printf("store the number %d into the Postgres database.\n", number)
	return nil
}

// 定义结构体
type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	if err := apiServer.numberStore.Put(123123); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The Numbers %+v\n", numbers)
}
