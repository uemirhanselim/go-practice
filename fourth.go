package main

import "fmt"

type WeaponType int

const (
	Axe WeaponType = iota
	Sword
	Dagger
	Bow
	Spear
	Hammer
)

func getDamage(weapon WeaponType) int {
	switch weapon {
	case Axe:
		return 100
	case Sword:
		return 90
	default:
		panic("Invalid weapon")
	}

}

func main() {
	fmt.Println(getDamage(Hammer))
}
