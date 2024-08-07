//go:build windows
// +build windows

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/rpc"
	"os"
	"strings"
	"syscall"

	"xpathlib/aardio/jsonrpc"

	"golang.org/x/net/html"

	//"github.com/JamesHovious/w32"
	"github.com/antchfx/htmlquery"
)

func xPathParse_(in string, expr string) (out string, err error) {
	// 拦截panic
	defer func() {
		if err := recover(); err != nil {
			out = ""
			err = errors.New(err.(string))
		}
	}()

	out = ""
	doc, err := htmlquery.Parse(strings.NewReader(in))
	if err != nil {
		return out, err
	}

	// 只拿到第一个匹配的
	node := htmlquery.FindOne(doc, expr)
	var b bytes.Buffer
	err = html.Render(&b, node)

	if err == nil {
		out = b.String()
	}

	return out, err
}

// 接口
type RpcBridge struct{}

// 接口参数
type RpcBridgeArgs struct {
	Html string
	Expr string
}

// 接口返回值
type RpcBridgeReply struct {
	Filter string
	Error  string
}

func (t *RpcBridge) XPathParse(args *RpcBridgeArgs, reply *RpcBridgeReply) error {
	// func xPathParse(in string, expr string) (out string, err error)
	out, err := xPathParse_(args.Html, args.Expr)
	if err != nil {
		reply.Error = err.Error()
		reply.Filter = err.Error()
	}
	if len(out) > 0 {
		reply.Filter = out
	}
	return err
}

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

// 重定向stderr
// https://stackoverflow.com/questions/34772012/capturing-panic-in-golang
func setStdHandle(stdhandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	err := setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
	if err != nil {
		//log.Fatalf("Failed to redirect stderr to file: %v", err)
		//w32.MessageBox(0, "Error!", "Redirect stderr to stdout error!", 0)
		os.Exit(-1)
		return
	}
	os.Stderr = f
}

func main() {

	// f, err := os.OpenFile("error.txt",
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	//log.Println(err)
	// 	return
	// }
	// defer f.Close()

	// 重定向Stderr到Stdout 避免panic弹窗
	redirectStderr(os.Stdout)
	////////////////////////////////

	server := rpc.NewServer()
	server.Register(new(RpcBridge))

	//运行 RPC 服务端
	//fmt.Printf("JSONRPC START RUN\n")
	jsonrpc.Run(server)

	////////////测试///////////////
	// https://www.cnblogs.com/zhaof/p/11346412.html
	resp, err := http.Get("https://golang.google.cn/dl/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		fmt.Printf("HTTP LENTH: %d\n", len(body))
	}

	//out, err := xPathParse(string(body), `//a`) //[contains(@class,"download")]
	//out, err := xPathParse(string(body), `//div[@class="download"]`)
	//if err != nil {
	//panic(err)
	//}
	//if len(out) > 0 {
	//	fmt.Println(out)
	//}

	// var rb RpcBridge
	// var rba RpcBridgeArgs
	// var rbr RpcBridgeReply

	// rba.Html = string(body)
	// rba.Expr = `//a`

	// rb.XPathParse(&rba, &rbr)
	//fmt.Println(rbr.Filter)
	/////////////////////////////////////
}
