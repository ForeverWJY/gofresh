@echo off

cd /d %~dp0 & cd & echo.

color 09
echo 3. ��ʽ�� Swagger API �ĵ�ע��...
swag fmt --generalInfo main.go & echo.

echo 4. ���� Swagger API ˵���ĵ�...
call swag init --parseVendor --generalInfo main.go
REM echo %errorlevel%
if %errorlevel% == 0 ( echo before build successfully ) else ( echo before build failed)
REM color 0f
REM echo ��ɣ� & echo.
REM if not "%1%" == "NoPause" (
  REM pause
REM )
