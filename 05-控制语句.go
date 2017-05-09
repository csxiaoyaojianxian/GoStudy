/*
* @Author: csxiaoyao
* @Date:   2017-04-23 16:36:54
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 01:16:40
 */
/**
 * 基础知识补充（指针、递增递减语句）、if判断语句、for循环语句、switch选择语句、跳转语句
 */
package main

import "fmt"

func main() {

	/*
		【指针】
		不能进行指针运算
		操作符 & 取变量地址，使用 * 通过指针间接访问目标对象
		默认值为 nil 而非 NULL

		【递增递减语句】
		++ 、 -- 作为语句而非表达式
	*/
	a := 1
	var p *int = &a
	fmt.Println(p)  // 0xc0420361d0
	fmt.Println(*p) // 1

	/*
		【判断语句if】
		支持一个初始化表达式（可以是并行方式）
		左大括号必须和条件语句或else在同一行，支持单行模式
	*/
	if a == 1 {
		fmt.Println("sun")
	}
	if a := 1; a > 1 {
		fmt.Println(a)
	}

	/*
		【switch】
		选择语句switch
		可以使用任何类型或表达式作为条件语句
		* 不需要写break，一旦条件符合自动终止
		如希望继续执行下一个case，需使用fallthrough语句
		支持一个初始化表达式（可以是并行方式），右侧需跟分号
		左大括号必须和条件语句在同一行
	*/
	a = 2
	switch a {
	case 0:
		fmt.Println("a=0")
		fallthrough
	case 1, 2, 3:
		fmt.Println("a=1 || a=2 || a=3")
	default:
		fmt.Println("none")
	}

	switch a := 1; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough // a>=0 a>=1
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("none")
	}

	a = 2
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough // a>=0 a>=1
	case a >= 1:
		fmt.Println("a>=1")
	}

	// Type Switch
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf(" x 是 int 型")
	case float64:
		fmt.Printf(" x 是 float64 型")
	case func(int) float64:
		fmt.Printf(" x 是 func(int) 型")
	case bool, string:
		fmt.Printf(" x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	/*
		【for循环】
		有三种形式
	*/
	// 形式1：【 for { } 】和 C 的 for(;;) 一样
	a = 1
	for {
		if a >= 3 {
			break
		}
		fmt.Println(a) // 1 2
		a++
	}
	// 形式2：【 for condition { } 】和 C 的 while 一样
	a = 1
	for a < 3 {
		fmt.Println(a) // 1 2
		a++
	}
	// 形式3：【 for init; condition; post { } 】和 C 语言的 for 一样
	for a = 1; a < 3; a++ {
		fmt.Println(a) // 1 2
	}
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	// for key, value := range oldMap {
	//     newMap[key] = value
	// }
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	// 第 0 位 x 的值 = 1
	// 第 1 位 x 的值 = 2
	// 第 2 位 x 的值 = 3
	// 第 3 位 x 的值 = 5
	// 第 4 位 x 的值 = 0
	// 第 5 位 x 的值 = 0

	/*
		【跳转语句 goto, break, continue】
		标签名区分大小写，若不使用会造成编译错误
		break与continue配合标签可用于多层循环的跳出
		goto是调整执行位置，与其它2个语句配合标签的结果并不相同
	*/
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
}
