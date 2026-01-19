package main

import "fmt"

func main() {
	numbers := []int{1, 2, 4}

	numbers = append(numbers, 2)

	fmt.Println(numbers)
}
