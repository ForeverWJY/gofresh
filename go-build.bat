:: ������ CGO_ENABLED=0
:: ������ GOOS=windows
:: ������ GOARCH=amd64
@echo off

cd /d %~dp0

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64

:: ���� GIN_MODE=release
SET GIN_MODE=release

:: build ѹ��
go build -ldflags "-w -s" -o gofresh.exe
REM go build -o gofresh.exe

pause