package main

import "fmt"

type UserStorer interface {
	GetAll() ([]string, error)
	Add(string) error
}

type InMemoryUserStore struct {
	users []string
}

func (i *InMemoryUserStore) GetAll() ([]string, error) {
	return i.users, nil
}

func (i *InMemoryUserStore) Add(name string) error {
	fmt.Println("Adding user", name)
	i.users = append(i.users, name)
	return nil
}

type UserService struct {
	userStore UserStorer
}

func main() {
	service := UserService{
		userStore: &InMemoryUserStore{},
	}
	fmt.Println(service.userStore.GetAll())
	service.userStore.Add("emirhan")
	fmt.Println(service.userStore.GetAll())
	service.userStore.Add("alihan")
	fmt.Println(service.userStore.GetAll())

}
