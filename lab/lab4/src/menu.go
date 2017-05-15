package main

import (
	"fmt"
	. "node"
	"os"
	// "reflect"
	"./singlechain"
	"strconv"
	"strings"
)

func showCmd(h *Node) {
	// 获取 cmd 数组
	cmd := singlechain.GetAll(h)
	// 循环遍历输出
	fmt.Println("******************** menu ********************")
	for i := 0; i < len(cmd); i++ {
		fmt.Println("cmd" + strconv.Itoa(i+1) + ": \t" + cmd[i].Cmd + "\t ---> " + cmd[i].Desc)
	}
	fmt.Println("**********************************************")
}

/**
 * 指令内容初始化
 */
var head [cmdLength + 1]Node

func initMenu(*head) {
	// 初始化链表
	head = singlechain.CreateLinkTable()

	h.Next = &head[1]
	head[0] = h
	head[1] = Node{"help", "This is a help cmd.", Help, &head[2]}
	head[2] = Node{"version", "menu program v2.0.", nil, &head[3]}
	head[3] = Node{"print", "Print the menu.go.", PrintCode, &head[4]}
	head[4] = Node{"quit", "Quit form menu.", Quit, nil}
}

singlechain.LinkTable *head = nil
func main() {
	// 初始化指令内容
	initMenu(head)

	var cmd string
	Help()
	for {
		fmt.Print("Input cmd or order ->")
		fmt.Scanf("%s\n", &cmd)
		if cmd != "" {
			// 默认判断输入的是序号，不能转换成整型则认为是 cmd
			order, _ := strconv.Atoi(cmd)
			var result *Node
			if order > 0 && order <= cmdLength {
				// 输入的是序号
				result = singlechain.GetLoc(&h, order)
			} else {
				// 输入的是指令
				result = FindCmd(&h, cmd)

			}
			// 执行
			if result != nil {
				if result.Handler != nil {
					result.Handler()
				} else {
					fmt.Println(result.Desc)
				}
			} else {
				fmt.Print("Error: unsupported command, you can use 'help' to list the available commands\n")
			}
		}
	}

	fmt.Println(singlechain.GetAll(&h))

}

// 根据 cmd 查找节点，FindCmd 不具有通用性，置于 menu.go
func FindCmd(h *Node, cmd string) *Node {
	if cmd == "" || h == nil {
		return nil
	}
	n := h
	for n.Next != nil {
		n = n.Next
		if strings.EqualFold(cmd, n.Cmd) {
			return n
		}
	}
	return nil
}

/**
 * 指令函数
 */
var Help = func() {
	showCmd(&h)
}

var Quit = func() {
	os.Exit(0)
}

var PrintCode = func() {
	userFile := "menu.go"
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
