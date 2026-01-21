package main

import "fmt"

func main() {
	maps := map[string]int{
		"ali":   1,
		"veli":  2,
		"ahmet": 3,
	}
	for _, value := range maps {
		fmt.Printf("the value %d", value)
	}

	name := "ali"

	switch name {
	case "ALI":
		fmt.Println("ali")
		break
	case "VELI":
		fmt.Println("veli")
		return
	case "AHMET":
		fmt.Println("ahmet")
	default:
		fmt.Println("unknown")
	}
}
