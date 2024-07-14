package main

import "fmt"

// 定义基本类型
var (
	code    int64   = 21398412934798
	count   int32   = 10
	product string  = "可乐"
	amount  float32 = 4.5
	total   float64 = 45
	unit    byte    = 0x2
	unicode rune    = 'Z'
)

// 自定义结构体
// 如果结构体中没有任何东西时，将不会占用任何内存
type Player struct {
	name         string
	health       int
	attachkPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func getHealth(player Player) int {
	return player.health
}

func main() {
	struct_demo()
	maps_demo()
	slice_demo()
}

// struct 练习
func struct_demo() {
	fmt.Println("=============================")
	fmt.Println("结构体练习：")

	player := Player{
		name:         "NPC",
		health:       23,
		attachkPower: 365.2,
	}

	fmt.Printf("this is the player: %v\n", player)
	fmt.Printf("this is the player: %+v\n", player)
	fmt.Printf("player.getHealth(): %d\n", player.getHealth())
	fmt.Printf("getHealth(): %d\n", getHealth(player))
}

// map 练习
func maps_demo() {
	fmt.Println("=============================")
	fmt.Println("map练习：")

	// 定义映射初始化内容
	users := map[string]int{
		"byrne": 26,
	}

	fmt.Printf("users: %+v\n", users)

	// 定义空映射
	employees := map[string]int{}
	// 添加内容
	employees["paige"] = 25

	fmt.Printf("employees: %+v\n", employees)

	// 使用make函数为map分配内存
	accounts := make(map[string]int)
	accounts["Admin"] = 123
	accounts["Normal"] = 456
	fmt.Printf("accounts: %+v\n", accounts)

	// 获取map中的值
	admin, ok := accounts["admin"]
	if !ok {
		fmt.Println("admin is not exists.")
	} else {
		fmt.Printf("admin: %s\n", admin)
	}

	// 删除map中的元素
	delete(accounts, "Normal")

	// 遍历map中的所有元素
	for k, v := range accounts {
		fmt.Printf("%s %d\n", k, v)
	}
}

// slice 练习
func slice_demo() {
	fmt.Println("=============================")
	fmt.Println("切片练习：")

	// 定义空切片
	slice_empty := []int{}
	fmt.Printf("slice_empty: %+v\n", slice_empty)

	// 定义切片并赋值
	numbers := []int{1, 2, 3}
	fmt.Printf("numbers: %+v\n", numbers)

	// 使用make为slice分配内存
	// slice 是动态增长的所以在make的时候需要指定初始大小
	slice := make([]int, 2)
	fmt.Println(slice)

	// 追加切片元素
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(slice)

	fmt.Println("=============================")
}

type Weapon string

func getWeapon(weapon Weapon) string {
	return string(weapon)
}
