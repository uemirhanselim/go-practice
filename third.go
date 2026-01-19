package main

import "fmt"

func main() {
	//users := map[string]int{}
	users := make(map[string]int)
	users["esu"] = 23

	age, ok := users["esu"]

	if !ok {
		fmt.Println("Bar does not exist")
	} else {
		fmt.Printf("this exist: %d\n", age)
	}

}
