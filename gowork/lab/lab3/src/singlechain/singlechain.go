/**
 * 单链表
 */
package singlechain

import (
	. "node"
)

// // 定义节点，h为头节点
// type Node struct {
// 	// Data int
// 	Cmd     string
// 	Desc    string
// 	Handler func()
// 	Next    *Node
// }

// 返回第一个节点
func GetFirst(h *Node) *Node {
	if h.Next == nil {
		return nil
	}
	return h.Next
}

// 返回最后一个节点
func GetLast(h *Node) *Node {
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
func GetLength(h *Node) int {
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
func Insert(h, d *Node, p int) bool {
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
func GetLoc(h *Node, p int) *Node {
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
func GetAll(h *Node) []Node {
	if h == nil {
		return nil
	}
	var allNode []Node
	n := h
	for n.Next != nil {
		n = n.Next
		allNode = append(allNode, *n)
	}
	return allNode
}
