@echo off
set GOROOT=D:\golang\go1.22.3
set GOBIN=%GOROOT%\bin
set CGO_ENABLED=0

rem 编译windows
set GOOS=windows
set GOARCH=386

set GOPATH=E:\go
set GO111MODULE=on
set GOPROXY=https://goproxy.io,direct
set GOSUMDB=off

@echo on
%GOBIN%\go.exe build -ldflags "-s -w" -v . 