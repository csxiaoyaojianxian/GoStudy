package tool

// 定义数据节点
type Node struct {
	// Data int
	Cmd     string
	Desc    string
	Handler func()
}