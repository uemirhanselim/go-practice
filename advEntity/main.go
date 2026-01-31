package main

import (
	"fmt"
	"strings"
)

type Putter interface {
	Put(id int, value any) error
}

type TransformFunc func(s string) string

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, transformFunc TransformFunc) string {
	return transformFunc(s)
}

type Storage interface {
	Putter
	Get(id int) (any, error)
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
