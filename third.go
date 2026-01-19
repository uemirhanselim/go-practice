package main

import "fmt"

func main() {
	//users := map[string]int{}
	users := make(map[string]int)
	users["esu"] = 23

	for k, v := range users {
		fmt.Printf("this is the key: %s and the value %d", k, v)
	}

	delete(users, "esu")
	fmt.Println("delete:", users["esu"])
}
