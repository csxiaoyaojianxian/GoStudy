/*
* @Author: sunshine
* @Date:   2017-05-10 02:00:38
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 12:23:13
 */

package main

import "fmt"

func main() {
	/*
		& 和 *
	*/
	var a int = 10
	var ip *int
	ip = &a
	var ptr *int
	fmt.Printf("%x %x %d\n", &a, ip, *ip) // c04200a230 c04200a230 10
	if ptr == nil {
		fmt.Printf("%x\n", ptr) // nil空指针，实际输出0
	}

	/*
		【指向数组的指针】
	*/
	// 指向数组的指针
	h := [...]int{19: 1}
	var i *[20]int = &h
	fmt.Println(i) // &[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

	// 使用new创建数组，返回指向数组的指针
	k := new([10]int)
	k[1] = 1
	fmt.Println(k) // &[0 1 0 0 0 0 0 0 0 0]

	/*
		【指针数组】
	*/
	// 指针数组
	x, y := 1, 2
	j := []*int{&x, &y}
	fmt.Println(*j[1], j) // 2 [0xc042012318 0xc042012320]

	/*
		【指向指针的指针变量，多级指针】
	*/
	var pptr **int
	ptr = &a
	pptr = &ptr
	fmt.Printf("变量 a = %d\n", a)                  // 10
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)          // 10
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr) // 10

	/*
		【指针作为函数参数】
	*/
	var b = 20
	swap(&a, &b)
	fmt.Printf("交换后 a 的值 : %d\n", a) // 20
	fmt.Printf("交换后 b 的值 : %d\n", b) // 10
	// 更快的方法
	a, b = b, a
	fmt.Printf("再次交换后 a 的值 : %d\n", a) // 10
	fmt.Printf("再次交换后 b 的值 : %d\n", b) // 20
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
