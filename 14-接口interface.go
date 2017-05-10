/*
* @Author: csxiaoyao
* @Date:   2017-04-25 00:06:47
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 23:02:15
 */
/**
 * 接口的定义与基本操作、嵌入接口、类型断言、空接口与 type switch、接口转换
 */
package main

import "fmt"

type USB interface {
	Name() string
	Connector
}
type Connector interface {
	Connect()
}
type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect:", pc.name)
}

type TVConnector struct {
	name string
}

func (tv TVConnector) Connect() {
	fmt.Println("Connect:", tv.name)
}

func main() {
	pc := PhoneConnector{"PhoneConnector"}
	// 声明接口
	var a Connector
	// 指定接口
	// a = new(PhoneConnector)
	// a = pc
	a = Connector(pc)

	// 实际调用，使用的是接口对象
	a.Connect()
	Disconnect(a)

	// tv := TVConnector{"TVConnector"}
	// var b USB
	// b = USB(tv) // 报错
	// b.Connect()
	// Disconnect(b)
}

// func Disconnect(usb USB) {
func Disconnect(usb interface{}) {
	// if pc, ok := usb.(PhoneConnector); ok {
	// 	fmt.Println("Disconnected:", pc.name)
	// 	return
	// }
	// fmt.Println("Unknown device.")

	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

/*
【接口interface】

接口是一个或多个方法签名的集合，只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示
接口只有方法声明，没有实现，没有数据字段
接口可以匿名嵌入其它接口，或嵌入到结构中
将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
只有当接口存储的类型和对象都为nil时，接口才等于nil
接口调用不会做receiver的自动转换
接口同样支持匿名字段方法
接口也可实现类似OOP中的多态
空接口可以作为任何类型数据的容器
*/

// 结果：
// Connect: PhoneConnector
// Disconnected: PhoneConnector
