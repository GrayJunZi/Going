# Going

### 一、安装

(1). 下载SDK

进入[官方下载地址](https://go.dev/dl/)进行SDK下载。

(2). 选择目标操作系统

这里使用`Ubuntu 24.02`系统作为示例。

(3). 下载途径

- `直接下载` - 直接根据操作系统，选择对应安装包进行下载。
    - 通过 `sftp username@ip` 进入linux中，并使用 `put <local_path> <remote_path>` 命令将文件上传至linux系统中。
- `命令行下载:`
    - 通过命令 `wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz` 进行下载。

(4). 解压`tar.gz`安装包

使用`tar`命令将压缩包进行解压。
```bash
tar -xvf go1.22.5.linux-amd64.tar.gz
```

将解压出来的go文件夹移动到`/usr/local`目录中。
```bash
sudo mv go /usr/local
```

(5). 设置环境变量

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

(6). 设置环境变量为国内代理

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### 二、变量与常量

- 全局变量
- 局部变量
- 常量
- 变量声明

在函数外声明的所有变量都称为全局变量，在函数内部声明的变量称为局部变量。

(1). 变量
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

(2). 常量
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