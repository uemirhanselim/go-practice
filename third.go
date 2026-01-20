package main

import "fmt"

type Weapon string

func getWeapon(weapon Weapon) string {
	return string(weapon)
}
func mainthird() {
	numbers := []int{1, 2, 4}

	numbers = append(numbers, 2)

	fmt.Println(numbers)
}
