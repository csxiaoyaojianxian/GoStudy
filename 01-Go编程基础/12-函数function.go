/*
* @Author: csxiaoyao
* @Date:   2017-04-23 22:09:37
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 01:39:25
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

func closure2() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 多返回值
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// panic
func G() {
	fmt.Println("G")
}
func H() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in H")
		}
	}()
	panic("Panic in H")
}
func I() {
	fmt.Println("I")
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

	// 【匿名函数、声明函数变量】
	b := func() {
		fmt.Println("function")
	}
	b()

	// 【闭包】
	f := closure(10)
	fmt.Println(f(1)) // 11
	// nextNumber 为一个函数，函数 i 为 0
	nextNumber := closure2()
	// 调用 nextNumber 函数，i 变量自增 1 并返回
	fmt.Println(nextNumber()) // 1
	fmt.Println(nextNumber()) // 2
	fmt.Println(nextNumber()) // 3
	// 创建新的函数 nextNumber1
	nextNumber1 := closure2()
	fmt.Println(nextNumber1()) // 1
	fmt.Println(nextNumber1()) // 2

	// 【defer】
	// 按照调用顺序的相反顺序执行，即使函数发生严重错误也会执行，支持匿名函数的调用
	// 常用于资源清理、文件关闭、解锁以及记录时间等操作
	// 通过与匿名函数配合可在return之后修改函数计算结果
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) // 3 3 3
		}()
	}

	// 【异常处理】
	// Go 没有异常机制，但有 panic/recover 模式来处理错误
	// Panic 可以在任何地方引发，但recover只有在defer调用的函数中有效
	G()
	H()
	I() // G、Recover in H、I

}
