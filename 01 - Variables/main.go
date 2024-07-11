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

// 声明常量
// 表示常数的惯用方法基本上总是小写的，如果需要导出常量则首字母需要大写
const versionName = "alpha"

const (
	key    = 4563
	Secret = "The Secret Key"
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

	fmt.Printf("VersionName: %s\n", versionName)
}
