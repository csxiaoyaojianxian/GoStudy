/*
* @Author: csxiaoyao
* @Date:   2017-04-23 22:09:37
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-23 22:37:15
 */
/**
 * 函数概述、函数的定义与使用、不定长变参、传递值类型和引用类型、匿名函数与闭包、defer 用法、panic 与 recover
 */
package main

import "fmt"

// 函数名，参数列表，返回值列表/单返回值类型/空
// func A(a int, b string) (int, string) {

// }
// func B(a, b, c int) int {

// }
func C() {

}
func D() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
	// return a,b,c
}

// 不定长参数
func E(b string, a ...int) {
	fmt.Println(a) // [1 2 3]
}

// 注意：不定长参数...为值传递，普通数组作为参数为引用传递
func F(a *int) { // 传递指针
	*a = 2
	fmt.Println(*a)
}

// 闭包
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	/*
	   【函数function】
	   不支持：嵌套、重载和默认参数
	   支持：无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包
	*/
	E("sun", 1, 2, 3)
	a := 1
	F(&a)
	fmt.Println(a) // 2

	// 【匿名函数】
	b := func() {
		fmt.Println("function")
	}
	b()
}
