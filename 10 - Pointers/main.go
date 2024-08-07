package main

import "fmt"

type Player struct {
	HP int
}

func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. NEW HP -> ", player.HP)
}

// 函数接收器
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. NEW HP -> ", p.HP)
}

func main() {
	player := &Player{
		HP: 100,
	}
	TakeDamage(player, 10)
	player.TakeDamage(10)
	fmt.Printf("%+v\n", player)
}
