package main

import "fmt"

type Entity struct {
	name    string
	id      string
	version string
	posx    int
	posy    int
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	entity := SpecialEntity{
		specialField: 100.0,
	}

	fmt.Printf("%+v\n", entity)
}
