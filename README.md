# XpathCheck

XpathCheck使用aardio和go编写，方便写爬虫时对html进行xpath提取校验，作为aardio跨进程jsonrpc调用go的示例项目

### 上手指南

###### 开发前的配置要求

1. go sdk v1.20(兼容win7)
2. aardio v36.27

###### **安装步骤**

1. 获取go开发环境 [go](https://golang.google.cn/dl/)
2. 获取aardio开发环境 [aardio](https://aardio.com/)
3. 进入go目录下修改build.bat

```bat
set GOROOT=你的go安装目录
set GOPATH=任意
```

4. 执行build.bat编译，成功后移动xpathlib.exe到dist目录
5. 启动aardio打开default.aproj，发布生成XpathCheck.exe


### 文件目录说明

```
XpathCheck
│  default.aproj
│  main.aardio
├─dist
│      XpathCheck.exe
│      xpathlib.exe
├─go
│  │  go.mod
│  │  go.sum
│  │  main.go
│  └─aardio
│      │  aardio.go
│      └─jsonrpc
│          │  jsonrpc.go
│          └─tcp
│                  tcp.go
```

### 使用到的框架

- [htmlquery](http://github.com/antchfx/htmlquery)
- [aardio/jsonrpc](http://aardio.com)


