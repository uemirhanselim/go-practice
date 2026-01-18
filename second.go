package main

import "fmt"

var (
	stringVar string = "adana"
	intVar    uint   = 1
	sae       rune
)

type Player struct {
	name   string
	health int
	ad     float64
}

func (player Player) getHealt2() int {
	return player.health
}

func getHealth(player Player) int {
	return player.health
}

func mainq() {
	player := Player{
		name:   "emirhan",
		health: 100,
		ad:     100.2,
	}

	fmt.Printf("This is the main player health: %d\n", player.getHealt2())
}
