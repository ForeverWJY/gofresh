:: 先设置 CGO_ENABLED=0
:: 再设置 GOOS=windows
:: 再设置 GOARCH=amd64
@echo off

cd /d %~dp0

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64

:: 设置 GIN_MODE=release
SET GIN_MODE=release

:: build 压缩
go build -ldflags "-w -s" -o gofresh.exe
REM go build -o gofresh.exe

pause