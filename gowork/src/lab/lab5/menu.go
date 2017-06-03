package main

import (
	"fmt"
	. "lab/lab5/tool"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

// 初始化常量，指令数量
const userFile = "menu.go"

// 定义全局变量
var initLt *LinkTable
var cmd []string

// 接口
var lt ILinkTable

/**
 * 指令内容初始化
 */
func initMenu(lt *LinkTable) {
	// 定义节点1
	var tempDataNode1 *Node = new(Node)
	//fmt.Printf("size=%d\n", unsafe.Sizeof(*tempDataNode1))
	var cmd *string = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode1)) + uintptr(unsafe.Sizeof(tempDataNode1))))
	*cmd = string("help")
	var desc *string = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode1)) + uintptr(unsafe.Sizeof(tempDataNode1)) + uintptr(unsafe.Sizeof(string("")))))
	*desc = string("This is a help cmd.")
	var handler *func() = (*func())(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode1)) + uintptr(unsafe.Sizeof(tempDataNode1)) + uintptr(unsafe.Sizeof(string(""))) + uintptr(unsafe.Sizeof(string("")))))
	*handler = Help
	// 添加节点1
	tempNode1 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode1))
	lt.AddNode(tempNode1)

	// 定义节点2
	var tempDataNode2 *Node = new(Node)
	cmd = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode2)) + uintptr(unsafe.Sizeof(tempDataNode2))))
	*cmd = string("version")
	desc = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode2)) + uintptr(unsafe.Sizeof(tempDataNode2)) + uintptr(unsafe.Sizeof(string("")))))
	*desc = string("menu program v2.8.")
	//handler = (*func())(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode2)) + uintptr(unsafe.Sizeof(tempDataNode2)) + uintptr(unsafe.Sizeof(string(""))) + uintptr(unsafe.Sizeof(string("")))))
	//*handler = nil
	// 添加节点2
	tempNode2 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode2))
	lt.AddNode(tempNode2)

	// 定义节点3
	var tempDataNode3 *Node = new(Node)
	cmd = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode3)) + uintptr(unsafe.Sizeof(tempDataNode3))))
	*cmd = string("print")
	desc = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode3)) + uintptr(unsafe.Sizeof(tempDataNode3)) + uintptr(unsafe.Sizeof(string("")))))
	*desc = string("Print the menu.go.")
	handler = (*func())(unsafe.Pointer(uintptr(unsafe.Pointer(tempDataNode3)) + uintptr(unsafe.Sizeof(tempDataNode3)) + uintptr(unsafe.Sizeof(string(""))) + uintptr(unsafe.Sizeof(string("")))))
	*handler = PrintCode
	// 添加节点3
	tempNode3 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode3))
	lt.AddNode(tempNode3)

	//tempDataNode1 = Node{"help", "This is a help cmd.", Help,nil}
	//tempNode1 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode1))
	//lt.AddNode(tempNode1)
	//
	//tempDataNode2 := Node{"version", "menu program v2.8.", nil,nil}
	//tempNode2 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode2))
	//lt.AddNode(tempNode2)

	//tempDataNode3 := Node{"print", "Print the menu.go.", PrintCode,nil}
	//tempNode3 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode3))
	//lt.AddNode(tempNode3)
	//
	//tempDataNode4 := Node{"addcmd", "Add a cmd.", AddCmd,nil}
	//tempNode4 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode4))
	//lt.AddNode(tempNode4)
	//
	//tempDataNode5 := Node{"delcmd", "Delete a cmd.", DelCmd,nil}
	//tempNode5 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode5))
	//lt.AddNode(tempNode5)
	//
	//tempDataNode6 := Node{"quit", "Quit form menu.", Quit,nil}
	//tempNode6 := (*LinkTableNode)(unsafe.Pointer(&tempDataNode6))
	//lt.AddNode(tempNode6)
}

func main() {

	// 初始化链表
	initLt = CreateLinkTable()
	// 初始化指令内容
	initMenu(initLt)

	lt = initLt

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
			if order > 0 && order <= lt.GetLength() {
				// 输入的是序号
				result = lt.GetLocNode(order)
			} else {
				// 输入的是指令
				result = lt.SearchNode(searchCondition, strings.ToLower(cmd[0]))
			}
			// 执行
			if result != nil {
				if (*Node)(unsafe.Pointer(result)).Handler != nil {
					(*Node)(unsafe.Pointer(result)).Handler()
				} else {
					fmt.Println((*Node)(unsafe.Pointer(result)).Desc)
				}
			} else {
				fmt.Print("Error: unsupported command, you can use 'help' to list the available commands\n")
			}
		}
	}
}

// 查找判定条件
func searchCondition(ltn *LinkTableNode, args string) bool {
	//if args == "" || ltn == nil {
	//	return false
	//}
	//if strings.EqualFold(args, (*Node)(unsafe.Pointer(ltn)).Cmd) {
	//	return true
	//}
	return false
}

// 读取一行数据
func ScanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	return string(b)
}

// 打印所有 cmd
func showCmd(lt *LinkTable) {
	if lt == nil {
		return
	}
	// 获取 cmd 数组
	var cmd []Node
	tempNode := lt.Head
	for tempNode != nil {
		//var n Node
		//n.Cmd =
		//cmd = append(cmd,)

		fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(tempNode)) + uintptr(8))))
		tempNode = tempNode.Next
	}


	//fmt.Println(cmd[0])



	// 循环遍历输出
	fmt.Println("******************** menu ********************")
	for i := 0; i < len(cmd); i++ {
		 //fmt.Println("cmd" + strconv.Itoa(i+1) + ": \t" + cmd[i].Cmd + "\t ---> " + cmd[i].Desc)
		//fmt.Println(*cmd[i])
	}
	fmt.Println("**********************************************")



}

/**
 * 指令函数
 */
// 打印菜单
var Help = func() {
	showCmd(initLt)
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
var AddCmd = func() {
	//length := len(cmd)
	//if length == 1 {
	//	fmt.Println("no arguments")
	//} else if length == 2 {
	//	tempDataNode := Node{cmd[1], "No Description", nil}
	//	tempNode := LinkTableNode{&tempDataNode, nil}
	//	lt.AddNode(&tempNode)
	//	Help()
	//} else if length == 3 {
	//	tempDataNode := Node{cmd[1], cmd[2], nil}
	//	tempNode := LinkTableNode{&tempDataNode, nil}
	//	lt.AddNode(&tempNode)
	//	Help()
	//} else if length == 4 {
	//	tempDataNode := Node{cmd[1], cmd[2], nil}
	//	tempNode := LinkTableNode{&tempDataNode, nil}
	//	p, _ := strconv.Atoi(cmd[3])
	//	lt.InsertLocNode(&tempNode, p)
	//	Help()
	//} else if length > 4 {
	//	fmt.Println("too many arguments")
	//}
}

// 删除指令
// 格式：
// cmd> delcmd printname
// cmd> delcmd 4
var DelCmd = func() {
	length := len(cmd)
	if length == 1 {
		fmt.Println("no arguments")
	} else if length == 2 {
		order, _ := strconv.Atoi(cmd[1])
		var result *LinkTableNode
		if order > 0 && order <= lt.GetLength() {
			// 输入的是序号
			lt.DeleteLocNode(order)
		} else {
			// 输入的是指令
			result = lt.SearchNode(searchCondition, cmd[1])
			lt.DelNode(result)
		}
		Help()
	} else if length > 2 {
		fmt.Println("too many arguments")
	}
}
