package tool

import (
	"fmt"
	"strconv"
	"strings"
)

// 定义全局变量
var initLt *LinkTable

const MAX_ARGC = 3 // 最大参数数量

// 接口
var lt ILinkTable

/**
 * 指令内容初始化
 */
func MenuConfig(cmd string, desc string, handler func(int,*[]string)) {
	// 初始化链表
	if initLt == nil {
		initLt = CreateLinkTable()
		tempDataNode1 := Node{"help", "This is a help cmd.", Help}
		tempNode1 := LinkTableNode{&tempDataNode1, nil}
		initLt.AddNode(&tempNode1)
	}
	tempDataNode2 := Node{cmd, desc, handler}
	tempNode2 := LinkTableNode{&tempDataNode2, nil}
	initLt.AddNode(&tempNode2)
	lt = initLt
}


/* menu program */
func ExecuteMenu() {
	/* cmd line begins */
	var (
		result int
		argLength int
	)
	Help(0,nil)
	var pcmd string
	var argv []string
	for {
		fmt.Print("Input cmd or order -> ")
		//fmt.Scanf("%s\n", &cmdStr)
		result,argLength = ScanLine(&pcmd, &argv)
		if result > 0 {
			var r *LinkTableNode
			if result == 1 {
				// 输入的是序号
				order,_ := strconv.Atoi(pcmd)
				r = lt.GetLocNode(order)
			} else if result == 2 {
				// 输入的是指令
				r = lt.SearchNode(searchCondition, strings.ToLower(pcmd))
			}
			// 执行
			if r != nil {
				if r.Data.Handler != nil {
					r.Data.Handler(argLength,&argv)
				} else {
					fmt.Println(r.Data.Desc)
				}
			} else {
				fmt.Print("Error: unsupported command, you can use 'help' to list the available commands.\n")
			}
		} else if result == -1 {
			fmt.Println("Error")
		} else if result == -2 {
			fmt.Println("Too many arguments.")
		}
	}

}

// 查找判定条件
func searchCondition(ltn *LinkTableNode, args string) bool {
	if args == "" || ltn == nil {
		return false
	}
	if strings.EqualFold(args, ltn.Data.Cmd) {
		return true
	}
	return false
}

// 读取一行数据
// 参数一：指令错误返回-1，参数溢出返回-2，输入的是序号返回1，指令返回2
// 参数二：argv长度
func ScanLine(pcmd *string, argv *[]string) (int,int) {
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
	cmd := strings.Split(string(b), " ")
	// 默认判断输入的是序号，不能转换成整型则认为是 cmd
	order, _ := strconv.Atoi(cmd[0])
	if cmd[0] == "" {
		return -1,0
	}
	*pcmd = cmd[0]
	argLength := len(cmd)-1
	if argLength > MAX_ARGC {
		return -2,0
	}
	for i:=0;i<argLength;i++{
		*argv = append(*argv, cmd[i+1])
	}
	if order > 0 && order <= lt.GetLength() {
		// 输入的是序号
		return 1,argLength
	} else if order > 0 {
		// 序号长度溢出
		return -1,0
	} else {
		// 输入的是指令
		return 2,argLength
	}
	return 0,0
}

// 打印所有 cmd
func showCmd(lt *LinkTable) {
	// 获取 cmd 数组
	cmd := lt.GetAllNode()
	// 循环遍历输出
	fmt.Println("******************** menu ********************")
	for i := 0; i < len(cmd); i++ {
		fmt.Println("cmd" + strconv.Itoa(i+1) + ": \t" + cmd[i].Data.Cmd + "\t ---> " + cmd[i].Data.Desc)
	}
	fmt.Println("**********************************************")
}

/**
 * 指令函数
 */
// 打印菜单
var Help = func(argc int, argv *[]string) {
	showCmd(initLt)
}

// 添加指令
// 格式:
// cmd> addcmd PrintName
// cmd> addcmd PrintName sunjianfeng
// cmd> addcmd PrintName sunjianfeng 3
var AddCmd = func(argc int, argv *[]string) {
	if argc == 0 {
		fmt.Println("no arguments")
	} else if argc == 1 {
		tempDataNode := Node{(*argv)[0], "No Description", nil}
		tempNode := LinkTableNode{&tempDataNode, nil}
		initLt.AddNode(&tempNode)
		lt = initLt
		Help(0,nil)
	} else if argc == 2 {
		tempDataNode := Node{(*argv)[0], (*argv)[1], nil}
		tempNode := LinkTableNode{&tempDataNode, nil}
		initLt.AddNode(&tempNode)
		lt = initLt
		Help(0,nil)
	} else if argc == 3 {
		tempDataNode := Node{(*argv)[0], (*argv)[1], nil}
		tempNode := LinkTableNode{&tempDataNode, nil}
		p, _ := strconv.Atoi((*argv)[2])
		lt.InsertLocNode(&tempNode, p)
		Help(0,nil)
	} else if argc > 3 {
		fmt.Println("too many arguments")
	}
}

// 删除指令
// 格式：
// cmd> delcmd printname
// cmd> delcmd 4
var DelCmd = func(argc int, argv *[]string) {
	if argc == 0 {
		fmt.Println("no arguments")
	} else if argc == 1 {
		order, _ := strconv.Atoi((*argv)[0])
		var result *LinkTableNode
		if order > 0 && order <= lt.GetLength() {
			// 输入的是序号
			lt.DeleteLocNode(order)
		} else {
			// 输入的是指令
			result = lt.SearchNode(searchCondition, (*argv)[0])
			lt.DelNode(result)
		}
		Help(0,nil)
	} else if argc > 1 {
		fmt.Println("too many arguments")
	}
}
