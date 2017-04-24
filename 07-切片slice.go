/*
* @Author: csxiaoyao
* @Date:   2017-04-23 20:49:57
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-23 21:31:20
 */
/**
 * slice 概述、创建 slice、reslice 概述、append()与slice、copy()与slice
 */
package main

import "fmt"

func main() {

	/*
	   【切片Slice】
	   其本身并不是数组，它指向底层的数组
	   使用len()获取元素个数，cap()获取容量
	   一般使用make()创建
	   如果多个slice指向相同底层数组，其中一个的值改变会影响全部
	*/
	// 声明
	// var slice1 []int
	slice2 := make([]int, 3, 10)          // 类型、长度、容量
	fmt.Println(len(slice2), cap(slice2)) // 3 10
	slice3 := make([]int, 3)              // 类型、长度
	fmt.Println(len(slice3), cap(slice3)) // 3 3
	// 【方法 slice】
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 := a[5:len(a)]
	s1 = a[5:]
	fmt.Println(s1) // [6 7 8 9 0]
	s2 := a[:5]
	fmt.Println(s2) // [1 2 3 4 5]
	s3 := a[2:5]
	fmt.Println(len(s3), cap(s3)) // 3 8
	// 【方法 Reslice】
	// reslice时索引以被slice的切片为准
	// 索引不可以超过被slice的切片的容量cap()值
	// 【方法 Append】
	// 如果最终长度未超过追加到slice的容量则返回原始slice
	// 如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
	slice4 := make([]int, 3, 6)
	fmt.Printf("%p\n", slice4) // 0xc0420462a0
	slice4 = append(slice4, 1, 2, 3)
	fmt.Printf("%v %p\n", slice4, slice4) // [0 0 0 1 2 3] 0xc0420462a0
	slice4 = append(slice4, 1, 2, 3)
	fmt.Printf("%v %p\n", slice4, slice4) // [0 0 0 1 2 3 1 2 3] 0xc042034060
	// 【方法 Copy】
	slice5 := []int{1, 2, 3, 4, 5, 6}
	slice6 := []int{7, 8, 9}
	copy(slice6, slice5)
	fmt.Println(slice6) // [1 2 3]
	slice7 := []int{1, 2, 3, 4, 5, 6}
	slice8 := []int{7, 8, 9}
	copy(slice7, slice8)
	fmt.Println(slice7) // [7 8 9 4 5 6]

	// 【数组的完整拷贝】
	slice9 := []int{1, 2, 3, 4, 5}
	slice10 := slice9[:]
	fmt.Println(slice10) // [1 2 3 4 5]
}
