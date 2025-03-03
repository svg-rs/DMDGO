@echo off
REM This script is used to build with go build .
REM It will build the project and output the binary to the current directory.

go build -o build/DMDGO.exe

if %errorlevel% neq 0 (
    echo Build failed.
    exit /b %errorlevel%
)
echo Build succeeded.
exit /b 0