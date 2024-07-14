package main

import "fmt"

type WeaponType int

const (
	Axe         WeaponType = 1
	Sword       WeaponType = 2
	Knife       WeaponType = 3
	WoodenStick WeaponType = 4
)

func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "Sword"
	case Sword:
		return "SWORD"
	case Knife:
		return "KNIFE"
	case WoodenStick:
		return "WOODENSTICK"
	default:
		panic("weapon does not exist")
	}
	return ""
}

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case Knife:
		return 40
	case WoodenStick:
		return 1
	default:
		panic("weapon does not exist")
	}
}

type RoleType int

const (
	Administrator RoleType = iota
	Normal
	Guest
)

func main() {
	fmt.Printf("damage of weapon: %s %d\n", Knife, getDamage(Knife))
}
