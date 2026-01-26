package main

import "fmt"

type Position struct {
	x, y int
}

func (p Position) Move(d int) {
	fmt.Printf("Moving %d units\n", d)
}

type Player struct {
	Position
}

type Color int

func (c Color) String() string {
	switch c {
	case colorBlue:
		return "BLUE"
	case colorRed:
		return "RED"
	case colorGreen:
		return "GREEN"
	case colorYellow:
		return "YELLOW"
	case colorPurple:
		return "PURPLE"
	default:
		panic("invalid color")
	}
}

const (
	colorBlue Color = iota
	colorRed
	colorGreen
	colorYellow
	colorPurple
)

func main() {
	p := Player{}
	p.Move(1000)
	fmt.Println(colorRed)
}
