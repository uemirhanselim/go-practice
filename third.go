package main

import "fmt"

func main() {
	//users := map[string]int{}
	users := make(map[string]int)
	users["esu"] = 23
	fmt.Printf("this is: %v\n", users)
}
