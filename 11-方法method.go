/*
* @Author: sunshine
* @Date:   2017-04-24 21:59:00
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-25 00:06:11
 */
/**
 * 方法的声明与使用、类型别名与方法、Method Value 与 Method Expression、方法名称冲突与字段访问权限
 */
package main

import "fmt"

/*
【方法method】
Receiver 可以是类型的值或者指针
不存在方法重载
可以使用值或指针来调用方法，编译器会自动完成转换
方法可以调用结构中的非公开字段
*/
type A struct {
	Name string
}

type B struct {
	Name string
}

type TZ int

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name) // AA

	b := B{}
	b.Print()
	fmt.Println(b.Name) //
}
func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A") // A
}

// func (a A) Print(b int) {
// 	fmt.Println("A") // 错误
// }
func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B") // B
}

func (a *TZ) Print() {
	fmt.Println("TZ")
}
