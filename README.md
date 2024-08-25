# Going

---

Full Time Go Dev

1. [x] 介绍
2. [x] Golang 入门
3. [ ] 掌握并发
4. [ ] 酒店预订项目 JSON API
5. [ ] 构建微服务
6. [ ] 找工作指南
7. [ ] 掌握核心区块链开发
8. [ ] 编程会话


# 一、Golang 入门

## 1. 安装

### (1). 下载SDK

进入[官方下载地址](https://go.dev/dl/)进行SDK下载。

### (2). 选择目标操作系统

这里使用`Ubuntu 24.02`系统作为示例。

### (3). 下载途径

- `直接下载` - 直接根据操作系统，选择对应安装包进行下载。
    - 通过 `sftp username@ip` 进入linux中，并使用 `put <local_path> <remote_path>` 命令将文件上传至linux系统中。
- `命令行下载:`
    - 通过命令 `wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz` 进行下载。

### (4). 解压`tar.gz`安装包

使用`tar`命令将压缩包进行解压。
```bash
tar -xvf go1.22.5.linux-amd64.tar.gz
```

将解压出来的go文件夹移动到`/usr/local`目录中。
```bash
sudo mv go /usr/local
```

### (5). 设置环境变量

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

### (6). 设置环境变量为国内代理

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 2. 变量与常量

- 全局变量
- 局部变量
- 常量
- 变量声明

在函数外声明的所有变量都称为全局变量，在函数内部声明的变量称为局部变量。

### (1). 变量
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

### (2). 常量
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

## 3. 内置类型与自定义类型

- 基本数据类型
- 结构体(structs)
- 集合(maps)
- 切片(slices)
- 数组(arrays)
- 自定义类型


### (1). 基本数据类型

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

### (2). 结构体

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

### (3). map 类型

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

### (4). 切片

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

### (5). 自定义类型

使用 `type` 来自定义类型
```go
type age int
type weapon string
```

## 4. 枚举

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

## 5. 控制结构

- for 循环
- range 循环
- break 退出循环

### (1). for 循环

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

### (2). range 循环

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

### (3). 退出循环

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

## 6. 接口


### (1). 使用 `interface` 关键字来定义接口
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

## 7. 包和模块

### (1). 创建模块

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

### (2). 定义包名

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

### (3). 访问性

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

### (4). Makefile

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

## 8. 高级类型

### (1). 结构体

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

### (2). 结构体的指针函数调用
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

### (3). 枚举

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

## 9. 高级接口与类型函数

### (1). 接口组合
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

### (2). 定义函数类型

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

## 10. 泛型


在定义的类型中使用 `[]` 定义泛型参数列表。
```go
package main

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 1)
	m1.Insert("bar", 2)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 9.99)
	m2.Insert(2, 89.23)
}
```

## 11. 指针

在go中，将某个值传入到函数中，都会在内存中得到一个该值的副本，对这个副本的所有操作都不会影响原始值。
```go
package main

import "fmt"

type Player struct {
	HP int
}

func TakeDamage(player Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. NEW HP -> ", player.HP)
}

func main() {
	player := Player{
		HP: 100,
	}
	TakeDamage(player, 10)
	fmt.Printf("%+v\n", player)
}
```

当函数接收一个指针的时候，将不会创建该值的副本，而是找到指针指向的内存空间，对该指针的所有操作都会影响到原始值，因为指针指向的内存地址为同一块区域。
```go
package main

import "fmt"

type Player struct {
	HP int
}

func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. NEW HP -> ", player.HP)
}

func main() {
	player := Player{
		HP: 100,
	}
	TakeDamage(&player, 10)
	fmt.Printf("%+v\n", player)
}
```

下面的使用方式与上面等效
```go
func main() {
	player := &Player{
		HP: 100,
	}
	TakeDamage(player, 10)
	fmt.Printf("%+v\n", player)
}
```

我们也可以定义函数接收器使用起来会更方便。
```go
// 函数接收器
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. NEW HP -> ", p.HP)
}

func main() {
	player := &Player{
		HP: 100,
	}
	player.TakeDamage(10)
	fmt.Printf("%+v\n", player)
}
```

## 12. 测试

### (1). 创建一个 `main.go` 文件，并添加一个函数。
```go
package main

type Player struct {
	name string
	hp   int
}

func calculateValues(x int, y int) int {
	return x + y
}

func main() {

}
```

添加测试代码，只需要创建一个和需要测试的文件同名并追加`_test`后缀的文件即可。例如 `main_test.go`。

添加测试方法命名通常以`Test`开头后面跟上要测试的方法名称。

```go
package main

import (
	"reflect"
	"testing"
)

func TestEqualPlayers(t *testing.T) {
	expected := Player{
		name: "Mark",
		hp:   100,
	}

	have := Player{
		name: "Alice",
		hp:   59,
	}

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v but got %+v", expected, have)
	}
}

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 6
	)

	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("excepted %d but have %d", expected, have)
	}
}

```

使用以下命令运行测试
```bash
go test ./...
```

要显示更详细的测试信息需要加 `-v`。
```bash
go test -v ./...
```

### (2). 我们可以指定测试某个方法。
```go
go test ./ -v -run TestEqualPlayers
```

### (3). 测试过的代码会将结果存入缓存中，可以通过`-count=1`的方式指定不使用缓存。
```
go test ./ -v -run TestEqualPlayers -count=1
```

# 二、Golang 并发

## 13. goroutines 和 channels 介绍

并发是Go语言中真正发光的地方。

### (1). 使用`go`关键字很容易就开启一个协程(goroutine)。
```go
package main

import (
	"time"
)

func main() {
	// 开启协程
	go fetchResource(1)
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
```

也可以调用匿名函数
```go
func main() {
	// 调用匿名函数
	go func() {
		result := fetchResource(1)
		fmt.Println(result)
	}()
}
```

### (2). 使用 `chan` 来声明一个 channel

```go
package main

import "fmt"

func main() {
	// 无缓冲通道
	resultch := make(chan string)

	// 有缓冲通道
	// resultch2 := make(chan string, 10)


	// 将数据写入 channel 中
	resultch <- "foo"

	// 将数据从 channel 中读取出来
	result := <-resultch

	fmt.Println(result)
}
```

`channel` 有两种类型，一种是无缓冲通道(unbuffered channel)，另一种是有缓冲通道(buffered channel)。

> Go语言中的 `channel` 无论是 无缓冲通道还是有缓冲通道 只要通道的内容已满 则将会阻塞。
>
> 无缓冲通道当写入一个数据时就会进行阻塞。当通道中没有数据时，读取通道将也会进行阻塞。 

## 14. 使用 Channel

### (1). 使用 `<-` 写入数据到channel中。

```go
package main

import "fmt"

func main() {
	// 缓冲长度最好是使用2的次方
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
}
```

### (2). 循环从channel中读取数据

注意如果此时不使用`close`将channel关闭将会在运行过程中产生死锁。
```go
package main

import "fmt"

func main() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	for msg := range msgch {
		fmt.Println("the msg is:", msg)
	}
}

```

### (3). 从channel获取数据时还可以额外获取一个状态

```go
package main

import "fmt"

func main() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	for {
		msg, ok := <-msgch
		if !ok {
			return
		}

		fmt.Println(msg)
	}
}
```

## 15. 控制流与异步

```go
package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 0 bytes
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop()
}

func (s *Server) quit() {
	s.quitch <- struct{}{}
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		}
	}
	fmt.Println("server is shutting down gracefully")
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message:", msg)
}

func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()
}
```

## 16. 互斥与原子值

Go语言中并发性是很重要的东西，这基本上是该语言被构建的原因。

### (1). 使用 `mutex` 来实现互斥
```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type State struct {
	mu      sync.Mutex
	count   int
	count32 int32
}

func (s *State) setCount(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count += i
}

func main() {
	state := &State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			state.setCount(i + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("%+v\n", state)
}

```


### (2). 使用 `atomic` 来实现原子操作

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type State struct {
	mu      sync.Mutex
	count int32
}

func (s *State) setCount(i int) {
	atomic.AddInt32(&s.count, int32(i))
}

func main() {
	state := &State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			state.setCount(i + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("%+v\n", state)
}

```

## 17. context 包

假设我们请求第三方的接口，如果该接口响应时间很长，我们不希望一直等待，而是有个超时机制，那么就可以使用`context`中已经定义好的超时机制来实现。

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "grayjunzi")
	userId, err := fetchUserData(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v -> %+v\n", time.Since(start), userId)
}

func fetchUserData(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	val := ctx.Value("username")
	fmt.Println("the value = ", val)

	type result struct {
		userId string
		err    error
	}
	resultch := make(chan result, 1)

	go func() {
		res, err := fetchData()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// Done
	// 1. -> context 已经超时
	// 2. -> context 已经被手动取消
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}
}

func fetchData() (string, error) {
	time.Sleep(time.Millisecond * 100)
	return "user id 1", nil
}
```

## 18. 实际例子

```go
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	Id       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)
	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)
	wg.Add(3)
	wg.Wait()
	close(respch)

	userProfile := &UserProfile{}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Hey, that was great",
		"Yeah Buddy",
		"Ow, I didnt know that",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}

	wg.Done()
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	respch <- Response{
		data: 14,
		err:  nil,
	}

	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	friendIds := []int{12, 43, 65, 87}
	respch <- Response{
		data: friendIds,
		err:  nil,
	}

	wg.Done()
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userProfile)
	fmt.Println("fetching te user profile took", time.Since(start))
}
```

## 19. 酒店预订项目

### (1). 项目概述

- 用户 -> 浏览酒店房间
- 管理员 -> 检查预约和预订
- 身份认证与鉴权 -> JWT
- 酒店 -> 增删改查的API
- 房间 -> 增删改查的API
- 脚本 -> 数据库管理、种子、迁移

### (2). 初始化项目

1. 初始化模块
```go
go mod init github.com/grayjunzi/hotel-reservation
```

2. 安装 gofiber
```go
go get github.com/gofiber/fiber/v2
```
3. 安装 MongoDB 驱动
```go
go get go.mongodb.org/mongo-driver/mongo
```

4. 创建Makefile简化命令
```makefile
build:
	go build -o bin/api

run: build
	./bin/api

test:
# 使用 @ 符号 隐藏输出
	@go test -v ./...
```

### (3). 添加路由

使用 fiber 来添加路由。
```go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/foo", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"msg": "i'm working on it."})
	})

	app.Listen(":5000")
}
```

可以使用`Group`声明一组路由。
```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/api"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/api/v1")

	v1.Get("/user", api.HandleGetUser)

	app.Listen(":5000")
}
```

使用`flag`从命令行中读取配置信息。
```go
package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()

	v1 := app.Group("/api/v1")

	v1.Get("/user", api.HandleGetUser)

	app.Listen(*listenAddr)
}
```

### (4). 创建MongoDB容器

```
docker run -d -p 27017:27017 --name mongodb mongodb/mongodb-community-server:latest
```

### (5). 安装加密库
```bash
go get golang.org/x/crypto/bcrypt
```