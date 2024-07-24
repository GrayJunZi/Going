# Going

## 一、安装

### 1. 下载SDK

进入[官方下载地址](https://go.dev/dl/)进行SDK下载。

### 2. 选择目标操作系统

这里使用`Ubuntu 24.02`系统作为示例。

### 3. 下载途径

- `直接下载` - 直接根据操作系统，选择对应安装包进行下载。
    - 通过 `sftp username@ip` 进入linux中，并使用 `put <local_path> <remote_path>` 命令将文件上传至linux系统中。
- `命令行下载:`
    - 通过命令 `wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz` 进行下载。

### 4. 解压`tar.gz`安装包

使用`tar`命令将压缩包进行解压。
```bash
tar -xvf go1.22.5.linux-amd64.tar.gz
```

将解压出来的go文件夹移动到`/usr/local`目录中。
```bash
sudo mv go /usr/local
```

### 5. 设置环境变量

在 `/etc/profile`文件中添加以下内容
```bash
export PATH=$PATH:/usr/local/go/bin
```

执行环境变量脚本内容
```bash
source /etc/profile
```

查看go所在位置
```bash
which go
```

查看go版本
```bash
go version
```

### 6. 设置环境变量为国内代理

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 二、变量与常量

- 全局变量
- 局部变量
- 常量
- 变量声明

在函数外声明的所有变量都称为全局变量，在函数内部声明的变量称为局部变量。

### 1. 变量
```go
package main

import "fmt"

// 全局变量
var name = "foo"

// 声明变量及类型并赋初始值
var firstName string = "Bar"

// 空字符串，go总是会初始化每一个变量如果没有赋值，将会自动分配为该类型的默认值
var lastName string

// 同时声明多个变量
var (
	a = 123
	b = "张三"
	c string
)

func main() {
	// 1. 仅声明变量
	var integer int

	// 2. 声明变量并赋值
	var float float32 = 10.25

	// infer: 自动推断类型并为变量赋值
	// 注意：使用:=的方式只能为新的变量赋值，已声明过的变量再使用这种方式赋值将会编译错误
	version := 1
	outVersion := "v1.0.1"

	fmt.Printf("interger: %d\n", integer)
	fmt.Printf("float: %.2f\n", float)
	fmt.Printf("version: %d\n", version)
	fmt.Printf("outVersion: %s\n", outVersion)
}
```

> 使用 `go run main.go` 命令运行go文件。

### 2. 常量
```go
package main

import "fmt"

// 声明常量
// 表示常数的惯用方法基本上总是小写的，如果需要导出常量则首字母需要大写
const versionName = "alpha"

// 同时声明多个常量
const (
	key    = 4563
	Secret = "The Secret Key"
)

func main() {
	fmt.Printf("VersionName: %s\n", versionName)
}
```

## 三、内置类型与自定义类型

- 基本数据类型
- 结构体(structs)
- 集合(maps)
- 切片(slices)
- 数组(arrays)
- 自定义类型


### 1. 基本数据类型

定义几种常见的基本数据类型
```go
var (
	code    int64   = 21398412934798
	count   int32   = 10
	product string  = "可乐"
	amount  float32 = 4.5
	total   float64 = 45
	unit    byte    = 0x2
	unicode rune    = 'Z'
)
```

### 2. 结构体

使用 `type` 以及 `struct` 关键字定义结构体类型，如果结构体中没有任何东西时，将不会占用任何内存。 
```go
// 定义结构体
type Player struct {
	name         string
	health       int
	attachkPower float64
}

// 定义 Player 的接收器方法
func (player Player) getHealth() int {
	return player.health
}

// 定义普通函数
func getHealth(player Player) int {
	return player.health
}

// struct 练习
func struct_demo() {
	fmt.Println("=============================")
	fmt.Println("结构体练习：")

	// 为结构体赋值
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
```

### 3. map 类型

通过 `map` 关键字来定义映射。
```go
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
```

### 4. 切片

使用 `[]int{}` 定义空的切片
```go
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
```

### 5. 自定义类型

使用 `type` 来自定义类型
```go
type age int
type weapon string
```

## 四、枚举

使用 `type` 定义类型后，结合 `const ()` 的方式将一组常量放在一起，可以组合成枚举的形式。
```go
type WeaponType int

const (
	Axe         WeaponType = 1
	Sword       WeaponType = 2
	Knife       WeaponType = 3
	WoodenStick WeaponType = 4
)
```

可以使用 `iota` 来进行简化操作，在第一个位置上为常量指定 iota，下面的常量中自动使用上面的类型并自动增长，初始值从0开始，然后每次加1。
```go
type RoleType int

const (
	Administrator RoleType = iota
	Normal
	Guest
)
```

## 五、控制结构

- for 循环
- range 循环
- break 退出循环

### 1. for 循环

使用 `for` 关键字进行指定次数循环。
```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```

也可以通过切片的长度，来实现遍历切片。
```go
numbers := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
for i := 0; i < len(numbers); i++ {
	fmt.Printf("numbers: %d\n", numbers[i])
}
```

### 2. range 循环

使用 `range` 关键字来遍历整个切片的内容，`i`循环的索引,`number`是元素值。
```go
for i, number := range numbers {
	fmt.Printf("range: index %d, number %d\n", i, number)
}
```

使用 `_` 来忽略索引，只获取值。
```go
for _, number := range numbers {
	fmt.Printf("range: ignore index, number %d\n", number)
}
```

### 3. 退出循环

使用 `break` 关键字来退出循环。
```go
users := map[string]int{
	"foo":   1,
	"bar":   2,
	"baz":   3,
	"alice": 4,
	"bob":   5,
}

for key, value := range users {
	if key == "alice" {
		break
	}
	fmt.Printf("key: %s value: %d\n", key, value)
}
```

## 六、接口

使用 `interface` 关键字来定义接口
```go
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}
```

实现接口，由于Go中采用鸭子类型的设计，所以不需要显示实现接口，只要将接口中定义的所有方法及属性都实现就被认为是该接口的实现。
```go
// 定义 MongoDB 实现
type MongoDBNumberStore struct {
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Printf("store the number %d into the MongoDB database.\n", number)
	return nil
}

// 定义 Postgres 实现
type PostgresNumberStore struct {
}

func (s PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (s PostgresNumberStore) Put(number int) error {
	fmt.Printf("store the number %d into the Postgres database.\n", number)
	return nil
}

```

> 鸭子类型：“如果它走起路来想鸭子，叫起来像鸭子，那么它就是鸭子。”
> 也就是说在Go中具有某个接口的全部方法和属性，就被认为是该接口的实现。


我们可以在结构体中定义某个接口的属性，然后在调用时传入该接口的具体实现。
```go
// 定义结构体
type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	if err := apiServer.numberStore.Put(123123); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The Numbers %+v\n", numbers)
}
```

## 七、包和模块

### 1. 创建模块

初始化模块
```go
go mod init myproject
```

编译项目
```go
go build -o ./myproject
```

运行项目
```go
./myproject
```

### 2. 定于包名

可以通过 `package` 自定义包名，但需要注意的是同一个文件夹下，只能声明同一个包名否则将报错。
```go
package types

type User struct {
	Username string
	Age      int
}
```

使用某个包下面的结构时可以使用`import`将包导入进来。
```go
package main

import (
	"fmt"
	"myproject/types"
)

func main() {
	user := types.User{
		Username: "James",
		Age:      getNumber(),
	}

	fmt.Printf("%s is %d years old.\n", user.Username, user.Age)
}
```

### 3. 访问性

定义的结构体、函数、属性等，如果名称首字母大写将被认为是公开的，可以被外部访问，如果首字母是小写则被认为是私有的，外部无法进行访问，只能内部才可以访问。

```go
package util

func GetUsername() string {
	return "James"
}

func GetNumber() int {
	return 84
}

```

### 4. Makefile

每次都要先编译项目然后再运行项目很麻烦，我们可以通过`Makefile`文件来简化操作。

```makefile
build:
	@go build -o myapp

run: build
	@./myapp
```

安装 `make`
```bash
sudo apt install make
sudo apt install make-guile
```

执行 `Makefile` 文件。
```bash
make run
```

## 八、高级类型

### 1. 结构体

go语言中没有继承操作，但是可以通过结构体的组合嵌套，来达到类似的目的。
```go
type Position struct {
	x, y int
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
}
```

### 2. 结构体的指针函数调用
```go
type Position struct {
	x, y int
}

func (p *Position) Move(val int) {
	fmt.Println("The position is moved by:", val)
}

type Player struct {
	Position
}

func main() {
	p := Player{}
	p.Move(1000)
}
```

### 3. 枚举

可以为枚举编写一个`String`函数，用于将枚举转换为字符串。

```go
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
```

## 九、高级接口与类型函数

### 1. 接口组合
```go
package main

type Putter interface {
	Put(id int, val any) error
}

type Storage interface {
	Putter
	Get(id int) (any, error)
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct{}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	storage Storage
}

func main() {
	s := &Server{
		storage: &FooStorage{},
	}
	s.storage.Get(1)
	s.storage.Put(1, "foo")
}
```

### 2. 定义函数类型

在函数定义中可以返回另一个函数
```go
package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("hello Sailor!", Uppercase))
	fmt.Println(transformString("hello Sailor!", Prefixer("FOO_")))
}
```