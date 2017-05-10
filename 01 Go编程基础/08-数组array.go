/*
* @Author: csxiaoyao
* @Date:   2017-04-23 18:46:01
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 12:12:03
 */
/**
 * array的定义、指向数组的指针和指针数组、数组之间的比较、使用new创建数组、多维数组、冒泡排序
 */
package main

import "fmt"

func main() {
	/*
	   【数组Array】
	   定义数组的格式：var <varName> [n]<type>，n>=0
	   数组长度是类型的一部分，因此具有不同长度的数组为不同类型
	   数组在Go中为值类型
	   数组之间可用==或!=比较，但不可使用<或>

	*/
	var a [2]string     // [ ]
	b := [2]int{1, 1}   // [1 1]
	c := [2]int{1}      // [1 0]
	d := [20]int{19: 1} // 索引 [0 0 … 0 1]
	// 忽略个数，根据初始化值数量判断，若出现索引，取最大索引
	e := [...]int{1, 2, 3, 4, 5}    // [1 2 3 4 5]
	f := [...]int{0: 1, 1: 2, 2: 3} // [1 2 3]
	g := [...]int{19: 1}            // 长度为20 [0 0 … 0 1]
	h := []int{2: 3}                // [0 0 3]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	/*
		【向函数传递数组】
	*/
	// void myFunction(param [10]int)
	// {
	//
	// }
	// void myFunction(param []int)
	// {
	//
	// }

	// 【多维数组】
	// var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
	i := [2][3]int{
		{1: 1},
		{2, 2, 2}}
	fmt.Println(i) // [[0 1 0] [2 2 2]]

	// 【冒泡排序】
	arr := [...]int{5, 2, 6, 3, 9}
	num := len(arr)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr) // [2 3 5 6 9]

}
