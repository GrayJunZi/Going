package main

import "fmt"

type Position struct {
	x, y int
}

func (p *Position) Move(val int) {
	fmt.Println("The position is moved by:", val)
}

type Player struct {
	Position
}

type Entity struct {
	id      string
	name    string
	version string

	Position
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	e := &SpecialEntity{
		specialField: 76.33,
		Entity: Entity{
			id:      "id 1",
			name:    "my entity",
			version: "version 1.0",
			Position: Position{
				x: 34,
				y: 87,
			},
		},
	}

	fmt.Printf("%+v\n", e)

	e.id = "new id"
	e.x = 56
	fmt.Printf("id: %s, x: %d\n", e.id, e.x)

	p := Player{}
	p.Move(1000)
}
