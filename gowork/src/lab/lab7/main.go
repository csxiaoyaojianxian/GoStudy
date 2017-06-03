package main

import (
	. "lab/lab7/tool"
	"os"
	"fmt"
)

// 初始化常量
const userFile = "tool/menu.go"

func main() {

	MenuConfig("version", "menu program v3.0.", nil)
	MenuConfig("print", "Print the menu.go.", PrintCode)
	MenuConfig("addcmd", "Add a cmd.", AddCmd)
	MenuConfig("delcmd", "Delete a cmd.", DelCmd)
	MenuConfig("quit", "Quit form menu.", Quit)
	ExecuteMenu()

}

// 退出
var Quit = func(argc int, argv *[]string) {
	os.Exit(0)
}

// 打印代码
var PrintCode = func(argc int, argv *[]string) {
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
