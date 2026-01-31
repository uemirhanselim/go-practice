package main

import "fmt"

type Storage interface {
	Get(id int) (any, error)
	Put(id int, value any) error
}

type UserStorage struct{}

func (s *UserStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *UserStorage) Put(id int, value any) error {
	return nil
}

func updateValue(id int, value any, storage Storage) error {
	return storage.Put(id, value)
}

type Server struct {
	storage Storage
}

func main() {
	server := &Server{
		storage: &UserStorage{},
	}

	server.storage.Put(1, "emirhan")
	value, err := server.storage.Get(1)

	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}

	fmt.Println("Value:", value)
}
