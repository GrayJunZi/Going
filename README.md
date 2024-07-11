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