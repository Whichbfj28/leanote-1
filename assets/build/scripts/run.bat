@echo off

cd..
set SCRIPTPATH=%cd%

: set GOPATH
set GOPATH="%SCRIPTPATH%\bin"

: run
if %processor_architecture%==x86 (
	"%SCRIPTPATH%\bin\leanote-windows-386.exe" -importPath github.com/coocn-cn/leanote
) else (
	"%SCRIPTPATH%\bin\leanote-windows-amd64.exe" -importPath github.com/coocn-cn/leanote
)
