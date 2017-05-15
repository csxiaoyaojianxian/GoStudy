package main

import (
	"fmt"
	. "node"
	"singlechain"
	"strconv"
	"strings"
)

var Help = func() {
	fmt.Println(111)
}

func main() {

	// 	static tDataNode head[] = {
	//     { "help","This is a help cmd.",Help,&head[1] },
	//     { "version","menu program v2.0.",NULL,&head[2] },
	//     { "quit","Quit form menu.",Quit,NULL }
	// };
	// 初始化头节点
	var h Node
	// 插入10个元素
	for i := 1; i <= 10; i++ {
		var d Node
		// d.Data = i
		d.Cmd = "cmd" + strconv.Itoa(i)
		d.Handler = Help

		singlechain.Insert(&h, &d, i)
		fmt.Println(singlechain.GetLoc(&h, i))
	}
	// fmt.Println(singlechain.GetLength(&h))
	// fmt.Println(singlechain.GetFirst(&h))
	// fmt.Println(singlechain.GetLast(&h))
	// fmt.Println(singlechain.GetLoc(&h, 6))

	singlechain.GetLoc(&h, 6).Handler()

	fmt.Println(FindCmd(&h, "cmd3"))
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
