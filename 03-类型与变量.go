/*
* @Author: csxiaoyao
* @Date:   2017-04-23 11:57:22
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-23 14:52:51
 */
/**
 * 基本类型、类型零值、类型别名、变量的声明与赋值、类型转换
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

// 类型别名
type (
	byte int8
	rune int32
	文本   string
)

func main() {

	/*
	   【Go基本类型】
	   布尔型：bool
	   	- 长度：1字节
	   	- 取值范围：true, false
	   	- 注意事项：不可以用数字代表true或false

	   整型：int/uint
	   	- 根据运行平台可能为32或64位

	   8位整型：int8/uint8
	   	- 长度：1字节
	   	- 取值范围：-128~127/0~255

	   字节型：byte（uint8别名）

	   16位整型：int16/uint16
	   	- 长度：2字节
	   	- 取值范围：-32768~32767/0~65535

	   32位整型：int32（rune）/uint32
	   	- 长度：4字节
	   	- 取值范围：-2^32/2~2^32/2-1/0~2^32-1

	   64位整型：int64/uint64
	   	- 长度：8字节
	   	- 取值范围：-2^64/2~2^64/2-1/0~2^64-1

	   浮点型：float32/float64
	   	- 长度：4/8字节
	   	- 小数位：精确到7/15小数位

	   复数：complex64/complex128
	   	- 长度：8/16字节
	   足够保存指针的 32 位或 64 位整数型：uintptr

	   其它值类型：
	   	- array、struct、string
	   引用类型：
	   	- slice、map、chan

	   接口类型：inteface
	   函数类型：func

	   【类型零值】
	   值类型默认值：0
	   布尔型默认值：false
	   字符串默认值：""

	*/

	var a []int
	fmt.Println(a) // []
	var b [1]bool
	fmt.Println(b) // [false]

	fmt.Println(math.MaxInt32) // 2147483647

	/*
		【类型别名】
		type newInt int，严格意义上讲是底层数据结构相同，不是别名，称为自定义类型，类型转换时仍需显式转换
		但 byte 和 rune 为 uint8 和 int32 的别名，可以相互转换
	*/
	var c 文本
	c = "字符串"
	fmt.Println(c) // 字符串

	/*
		【单个变量声明与赋值】
		变量的声明格式：var <变量名称> <变量类型>
		变量的赋值格式：<变量名称> = <表达式>
		声明的同时赋值：var <变量名称> [变量类型] = <表达式>
	*/
	var d int
	d = 1
	fmt.Println(d)
	var e = 1
	fmt.Println(e)
	f := 1
	fmt.Println(f)

	/*
		【多个变量声明与赋值】
		全局变量声明可使用 var()
		全局变量声明不可省略 var，但可使用并行方式
		所有变量都可使用类型推断
		局部变量不可使用 var() 的方式简写，只能使用并行方式
	*/
	var g, h, i int = 1, 2, 3
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	// 使用并行方式以及进行类型判断
	var j, k, l = 1, 2, 3
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	var (
		m    = 1
		n, o = 2, 3
	)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	// 冒号不能用于全局变量
	// 空白符号
	p, _, q := 1, 2, 3
	fmt.Println(p)
	fmt.Println(q)

	/*
		【类型转换】
		Go不存在隐式转换，所有类型转换必须显式声明
		转换只能发生在两种相互兼容的类型之间
		<ValueA> [:]= <TypeOfValueA>(<ValueB>)
	*/
	var r float32 = 1.1
	s := int(r)
	fmt.Println(s)
	// var t bool
	// t = bool(s)
	// fmt.Println(t)
	// var u string = "11"
	// v := int(u)
	// fmt.Println(v)

	// string() 表示将数据转换成文本格式
	var w int = 65
	x := string(w)
	fmt.Println(x) // A

	// strconv
	y := strconv.Itoa(w)
	fmt.Println(y) // 65
	z, _ := strconv.Atoi(y)
	fmt.Println(z) // 65

}
