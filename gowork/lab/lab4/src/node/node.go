package node

// 定义节点，h为头节点
type Node struct {
	// Data int
	Cmd     string
	Desc    string
	Handler func()
	Next    *Node
}
