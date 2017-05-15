/*
* @Author: csxiaoyao
* @Date:   2017-04-23 14:37:43
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-08 19:30:59
 */
/**
 * 常量的定义、常量的枚举、运算符
 */
package main

import "fmt"

const a int = 1
const b = 'A'
const (
	c       = a
	d       = c + 1
	e       = d + 2
	f, g, h = 1, "2", '3'
)
const i, j, k = 1, "2", '3'
const (
	l = 1
	m
	n
)

var temp string = "sun"

const (
	o = "sun"
	// p = len(temp)
	p = len(o)
)
const (
	q, r = 1, "sun"
	s, t
)

const (
	u  = 'A'
	v  // 'A'
	v1 = 100
	v2        // 100
	w  = iota // 2
	x         // 3
)

// 星期枚举
const (
	// 第一个常量不可省略表达式
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 计算机存储单位枚举
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {

	/*
		【常量的定义】
		常量的值在编译时就已经确定
		常量的定义格式与变量基本相同
		等号右侧必须是常量或者【常量表达式】
		常量表达式中的函数必须是【内置函数】

		【常量的初始化规则与枚举】
		在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
		使用相同的表达式不代表具有相同的值
		iota是常量的计数器，从0开始，组中每定义1个常量自动递增1

	*/
	fmt.Println(e) // 4
	fmt.Println(m) // 1
	fmt.Println(n) // 1
	fmt.Println(p) // 3
	fmt.Println(s) // 1
	fmt.Println(t) // sun

	fmt.Println(v) // 65
	fmt.Println(w) // 2
	fmt.Println(x) // 3

	fmt.Println(Friday) // 5

	fmt.Println(GB) // 1.073741824e+09

	/*
		【运算符】

		算术运算符 + - * / % ++ --
		关系运算符 == != > < >= <=
		逻辑运算符 && || !
		位运算符   & | ^ << >>
		赋值运算符 = += -= *= /= %= <<= >>= &= ^= |=
		其他运算符
			&	返回变量存储地址
			*	指针变量


		优先级（从高到低）
		^	!						（一元运算符）
		*	/	%	<<	>>	&	&^
		+	-	|	^				（二元运算符）
		==	!=	<	<=	>=	>
		<-							（专门用于channel）
		&&
		||

	*/
}
