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