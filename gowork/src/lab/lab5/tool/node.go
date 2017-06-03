package tool

// 定义数据节点
type Node struct {
	// Data int
	Next *Node     // 8 byte
	Cmd     string // 16 byte
	Desc    string // 16 byte
	Handler func() // 8 byte
}