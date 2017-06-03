/**
 * 单链表
 */
package tool

// 定义链表接口
type ILinkTable interface {
	//DeleteLinkTable()
	AddNode(ltn *LinkTableNode) bool
	DelNode(ltn *LinkTableNode) bool
	GetHead() *LinkTableNode
	GetTail() *LinkTableNode
	GetLength() int
	GetNextNode(ltn *LinkTableNode) *LinkTableNode
	InsertLocNode(ltn *LinkTableNode, p int) bool
	DeleteLocNode(p int) bool
	GetLocNode(p int) *LinkTableNode
	//GetAllNode() []*LinkTableNode
	SearchNode(condition func(*LinkTableNode, string) bool, args string) *LinkTableNode
}

// 定义节点，分为数据区和指针区
type LinkTableNode struct {
	//Data *Node
	Next *LinkTableNode
}

// 定义链表
type LinkTable struct {
	Head      *LinkTableNode
	Tail      *LinkTableNode
	SumOfNode int
}

// 创建链表
func CreateLinkTable() *LinkTable {
	var linkTable LinkTable
	linkTable.Head = nil
	linkTable.Tail = nil
	linkTable.SumOfNode = 0
	return &linkTable
}

// 删除链表，GO自动GC
// func (lt *LinkTable) DeleteLinkTable() bool {
// 	if lt == nil {
// 		return false
// 	}
//	tempNode := lt.Head
// 	for tempNode != nil {
//		freeNode := tempNode
// 		tempNode = tempNode.Next
//		if freeNode != nil {
//			freeNode = nil
//		}
// 	}
//	lt.Head = nil
//	lt.Tail = nil
//	lt.SumOfNode = 0
// 	lt = nil
// 	return true
// }

// 追加节点
func (lt *LinkTable) AddNode(ltn *LinkTableNode) bool{
	if lt == nil || ltn == nil {
		return false
	}

	if lt.Head == nil {
		// 链表为空
		lt.Head = ltn
		lt.Tail = ltn
		lt.SumOfNode = 1
	} else {
		// 链表不为空
		lt.Tail.Next = ltn
		lt.Tail = ltn
		lt.SumOfNode++
	}
	return true
}

// 删除节点
func (lt *LinkTable) DelNode(ltn *LinkTableNode) bool {
	if lt == nil || ltn  == nil {
		return false
	}
	// 删除头节点
	if lt.Head == ltn{
		lt.Head = lt.Head.Next
		lt.SumOfNode--
		// 删除头节点后可能链表为空，需要修改tail
		if lt.SumOfNode == 0 {
			lt.Tail = nil
		}
		return true
	}
	// 查找节点并删除
	tempNode := lt.Head
	for tempNode != nil {
		if tempNode.Next == ltn {
			tempNode.Next = ltn.Next
			// 如果删除的是tail，修改tail
			if lt.Tail == ltn {
				lt.Tail = tempNode
			}
			lt.SumOfNode--
			return true
		}
		tempNode = tempNode.Next
	}
	return false
}

// 返回头节点
func (lt *LinkTable) GetHead() *LinkTableNode {
	if lt == nil || lt.Head == nil{
		return nil
	}
	return lt.Head
}

// 返回尾节点
func (lt *LinkTable) GetTail() *LinkTableNode {
	if lt == nil || lt.Tail == nil{
		return nil
	}
	return lt.Tail
}

// 获取长度
func (lt *LinkTable) GetLength() int {
	return lt.SumOfNode
}

// 返回下一个节点
func (lt *LinkTable) GetNextNode(ltn *LinkTableNode) *LinkTableNode{
	if lt == nil || ltn == nil {
		return nil
	}
	tempNode := lt.Head
	for tempNode != nil {
		if ltn == tempNode {
			return ltn.Next
		}
		tempNode = tempNode.Next
	}
	return nil
}

// 指定位置插入节点
func (lt *LinkTable) InsertLocNode(ltn *LinkTableNode, p int) bool {
	if lt == nil || ltn == nil || p < 0 || p > lt.GetLength()+1 {
		return false
	}

	if lt.SumOfNode == 0 && p == 1{
		// 空链表,插入到第一位
		lt.Head = ltn
		lt.Tail = ltn
		lt.SumOfNode = 1
	}else if p == 1{
		// 非空链表,插入到第一位
		ltn.Next = lt.Head
		lt.Head = ltn
		lt.SumOfNode++
	}else if p == lt.GetLength()+1 {
		// 非空链表，插入到最后一位
		ltn.Next = nil
		lt.Tail.Next = ltn
		lt.Tail = ltn
		lt.SumOfNode++

	}else{
		tempNode := lt.GetLocNode(p-1)
		if tempNode == nil {
			return false
		}
		ltn.Next = tempNode.Next
		tempNode.Next = ltn
		lt.SumOfNode++
	}
	return false
}

// 指定位置删除节点
func (lt *LinkTable) DeleteLocNode(p int) bool {
	if lt == nil || p < 0 || p > lt.GetLength() || lt.SumOfNode <= 0 {
		return false
	}

	if lt.SumOfNode == 1 && p == 1{
		// 只有一个节点
		lt.Head = nil
		lt.Tail = nil
		lt.SumOfNode = 0
	}else if p == 1{
		// 删除头节点
		lt.Head = lt.Head.Next
		lt.SumOfNode--
	}else if p == lt.GetLength() {
		// 删除尾节点
		lt.Tail = lt.GetLocNode(p-1)
		lt.SumOfNode--
	}else{
		tempNode := lt.GetLocNode(p-1)
		if tempNode == nil {
			return false
		}
		tempNode.Next = tempNode.Next.Next
		lt.SumOfNode--
	}
	return false
}

// 指定位置取出节点
func (lt *LinkTable) GetLocNode(p int) *LinkTableNode {
	if lt == nil || p < 0 || p > lt.GetLength() {
		return nil
	}
	var i = 0
	tempNode := lt.Head
	for tempNode != nil {
		i++
		if i == p {
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}

// 获取所有节点
//func (lt *LinkTable) GetAllNode() []*LinkTableNode {
//	if lt == nil {
//		return nil
//	}
//	var allNode []*LinkTableNode
//	tempNode := lt.Head
//	for tempNode != nil {
//		allNode = append(allNode, tempNode)
//		tempNode = tempNode.Next
//	}
//	return allNode
//}


// 利用回调函数查找节点
func (lt *LinkTable) SearchNode(condition func(*LinkTableNode, string) bool, args string) *LinkTableNode {
	if args == "" || lt == nil {
		return nil
	}
	tempNode := lt.Head
	for tempNode != nil {
		if condition(tempNode, args) == true {
			return tempNode
		}
		tempNode = tempNode.Next
	}
	return nil
}