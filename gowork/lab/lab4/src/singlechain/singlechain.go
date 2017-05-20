/**
 * 单链表
 */
package singlechain

// 定义节点
type LinkTableNode struct {
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

// 删除链表
// func DeleteLinkTable(linkTable *LinkTable) bool {
// 	if linkTable == nil {
// 		return false
// 	}
// 	for linkTable.Head != nil {
// 		tmp := linkTable.Head
// 		linkTable.Head = linkTable.Head.Next
// 		tmp = nil
// 	}
// 	linkTable = nil
// 	return true
// }

// 返回第一个节点
func GetFirst(h *LinkTableNode) *LinkTableNode {
	if h.Next == nil {
		return nil
	}
	return h.Next
}

// 返回最后一个节点
func GetLast(h *LinkTableNode) *LinkTableNode {
	if h.Next == nil {
		return nil
	}
	i := h
	for i.Next != nil {
		i = i.Next
		if i.Next == nil {
			return i
		}
	}
	return nil
}

// 获取长度
func GetLength(h *LinkTableNode) int {
	var i int = 0
	n := h
	for n.Next != nil {
		i++
		n = n.Next
	}
	return i
}

// 插入节点
// h:head,d:node,p:place，从0开始计算
func Insert(h, d *LinkTableNode, p int) bool {
	// 空链表插入到头节点后
	if h.Next == nil {
		h.Next = d
		return true
	}
	i := 0
	n := h
	for n.Next != nil {
		// i 表示下一个节点
		i++
		// 插入在下下个节点位置
		if i == p {
			if n.Next.Next == nil {
				n.Next = d
				return true
			} else {
				d.Next = n.Next
				n.Next = d.Next
				return true
			}
		}
		n = n.Next
		if n.Next == nil {
			n.Next = d
			return true
		}
	}
	return false
}

// 取出指定节点
func GetLoc(h *LinkTableNode, p int) *LinkTableNode {
	if p < 0 || p > GetLength(h) {
		return nil
	}
	var i int = 0
	n := h
	for n.Next != nil {
		i++
		n = n.Next
		if i == p {
			return n
		}
	}
	return nil
}

// 显示所有节点
func GetAll(h *LinkTableNode) []LinkTableNode {
	if h == nil {
		return nil
	}
	var allNode []LinkTableNode
	n := h
	for n.Next != nil {
		n = n.Next
		allNode = append(allNode, *n)
	}
	return allNode
}
