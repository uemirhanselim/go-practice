package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDBNumberStore struct {
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("Putting number", number)
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func mainInterfaces() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	fmt.Println(apiServer.numberStore.GetAll())
}
