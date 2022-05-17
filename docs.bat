@echo off

cd /d %~dp0 & cd & echo.

color 09
echo 3. 格式化 Swagger API 文档注释...
swag fmt --generalInfo main.go & echo.

echo 4. 生成 Swagger API 说明文档...
call swag init --parseVendor --generalInfo main.go
REM echo %errorlevel%
if %errorlevel% == 0 ( echo before build successfully ) else ( echo before build failed)
REM color 0f
REM echo 完成！ & echo.
REM if not "%1%" == "NoPause" (
  REM pause
REM )
