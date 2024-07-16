package main

import "fmt"

type Color int

const (
	ColorRed Color = iota
	ColorBlue
	ColorGreen
	ColorBlack
	ColorYellow
)

func (c Color) String() string {
	switch c {
	case ColorRed:
		return "RED"
	case ColorBlue:
		return "BLUE"
	case ColorGreen:
		return "GREEN"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	default:
		panic("invalid color given.")
	}
}

func main() {
	fmt.Println("this red color is:", ColorRed)
	fmt.Println("this black color is:", ColorBlack)
}
