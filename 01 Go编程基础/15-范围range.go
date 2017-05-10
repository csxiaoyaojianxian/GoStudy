/*
* @Author: csxiaoyao
* @Date:   2017-05-10 13:44:38
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 13:56:12
 */

package main

import "fmt"

func main() {
	// range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)、集合(map)元素
	// 数组、切片返回元素索引值，集合中返回 key-value 对 key 值

	// range 迭代数组、slice，将传入 index 和 value
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // sum: 9

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) // index: 1
		}
	}

	// range 迭代 map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // a -> apple 、b -> banana
	}

	// range 枚举 Unicode 字符串，传入 index 和 value(Unicode值)
	for i, c := range "go" {
		fmt.Println(i, c) // 0 103 、1 111
	}
}
