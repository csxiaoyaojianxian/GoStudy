package main

import (
	"fmt"
	. "lab/lab4/tool"
	"os"
	"strings"
	"strconv"
)

// 初始化常量，指令数量
const userFile = "menu.go"

// 定义全局变量
var lt *LinkTable
var cmd []string

/**
 * 指令内容初始化
 */
func initMenu(lt *LinkTable) {
	tempDataNode1 := Node{"help", "This is a help cmd.", Help}
	tempNode1 := LinkTableNode{&tempDataNode1, nil}
	AddNode(lt,&tempNode1)

	tempDataNode2 := Node{"version", "menu program v2.0.", nil}
	tempNode2 := LinkTableNode{&tempDataNode2, nil}
	AddNode(lt,&tempNode2)

	tempDataNode3 := Node{"print", "Print the menu.go.", PrintCode}
	tempNode3 := LinkTableNode{&tempDataNode3, nil}
	AddNode(lt,&tempNode3)

	tempDataNode4 := Node{"addcmd", "Add a cmd.", AddCmd}
	tempNode4 := LinkTableNode{&tempDataNode4, nil}
	AddNode(lt,&tempNode4)

	tempDataNode5 := Node{"delcmd", "Delete a cmd.", DelCmd}
	tempNode5 := LinkTableNode{&tempDataNode5, nil}
	AddNode(lt,&tempNode5)

	tempDataNode6 := Node{"quit", "Quit form menu.", Quit}
	tempNode6 := LinkTableNode{&tempDataNode6, nil}
	AddNode(lt,&tempNode6)
}


func main() {
	// 初始化链表
	lt = CreateLinkTable()
	// 初始化指令内容
	initMenu(lt)

	var cmdStr string
	Help()
	for {
		fmt.Print("Input cmd or order -> ")
		//fmt.Scanf("%s\n", &cmdStr)
		cmdStr = ScanLine()
		if cmdStr != "" {
			cmd = strings.Split(cmdStr, " ")
			// 默认判断输入的是序号，不能转换成整型则认为是 cmd
			order, _ := strconv.Atoi(cmd[0])
			var result *LinkTableNode
			if order > 0 && order <= lt.SumOfNode {
				// 输入的是序号
				result = GetLocNode(lt, order)
			} else {
				// 输入的是指令
				result = FindCmd(lt, strings.ToLower(cmd[0]))
			}
			// 执行
			if result != nil {
				if result.Data.Handler != nil {
					result.Data.Handler()
				} else {
					fmt.Println(result.Data.Desc)
				}
			} else {
				fmt.Print("Error: unsupported command, you can use 'help' to list the available commands\n")
			}
		}
	}
}

// 读取一行数据
func ScanLine() string {
	var c byte
	var err error
	var b []byte
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break;
		}
	}
	return string(b)
}

// 打印所有 cmd
func showCmd(lt *LinkTable) {
	// 获取 cmd 数组
	cmd := GetAllNode(lt)
	// 循环遍历输出
	fmt.Println("******************** menu ********************")
	for i := 0; i < len(cmd); i++ {
		fmt.Println("cmd" + strconv.Itoa(i+1) + ": \t" + cmd[i].Data.Cmd + "\t ---> " + cmd[i].Data.Desc)
	}
	fmt.Println("**********************************************")
}

// 根据 cmd 查找节点，FindCmd 不具有通用性，置于 menu.go
func FindCmd(lt *LinkTable, cmd string) *LinkTableNode {
	if cmd == "" || lt == nil {
		return nil
	}
	tempNode := lt.Head
	for tempNode != nil {
		if strings.EqualFold(cmd, tempNode.Data.Cmd) {
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}

/**
 * 指令函数
 */
// 打印菜单
var Help = func() {
	showCmd(lt)
}
// 退出
var Quit = func() {
	os.Exit(0)
}
// 打印代码
var PrintCode = func() {
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
// 添加指令
// 格式:
// cmd> addcmd PrintName
// cmd> addcmd PrintName sunjianfeng
// cmd> addcmd PrintName sunjianfeng 3
var AddCmd = func(){
	length := len(cmd)
	if length == 1 {
		fmt.Println("no arguments")
	} else if length == 2 {
		tempDataNode := Node{cmd[1], "No Description", nil}
		tempNode := LinkTableNode{&tempDataNode, nil}
		AddNode(lt,&tempNode)
		Help()
	} else if length == 3 {
		tempDataNode := Node{cmd[1], cmd[2], nil}
		tempNode := LinkTableNode{&tempDataNode, nil}
		AddNode(lt,&tempNode)
		Help()
	} else if length == 4 {
		tempDataNode := Node{cmd[1], cmd[2], nil}
		tempNode := LinkTableNode{&tempDataNode, nil}
		p,_ := strconv.Atoi(cmd[3])
		InsertLocNode(lt,&tempNode,p)
		Help()
	} else if length > 4{
		fmt.Println("too many arguments")
	}
}
// 删除指令
// 格式：
// cmd> delcmd printname
// cmd> delcmd 4
var DelCmd = func(){
	length := len(cmd)
	if length == 1 {
		fmt.Println("no arguments")
	} else if length == 2 {
		order, _ := strconv.Atoi(cmd[1])
		var result *LinkTableNode
		if order > 0 && order <= lt.SumOfNode {
			// 输入的是序号
			DeleteLocNode(lt,order)
		} else {
			// 输入的是指令
			result = FindCmd(lt,cmd[1])
			DelNode(lt,result)
		}
		Help()
	} else if length > 2 {
		fmt.Println("too many arguments")
	}
}