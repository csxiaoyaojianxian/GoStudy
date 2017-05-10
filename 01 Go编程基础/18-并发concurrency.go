/*
* @Author: csxiaoyao
* @Date:   2017-04-25 00:57:41
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 23:38:44
 */

package main

import (
	"fmt"
	// "time"
)

func main() {

	/**
	 * 并发concurrency
	 */
	// 并发不是并行：Concurrency Is Not Parallelism
	// Go 可以设置使用核数，以发挥多核计算机的能力
	// Goroutine 奉行通过通信来共享内存，而不是共享内存来通信

	// 【Channel】
	// Channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
	// 通过 make 创建，close 关闭
	// Channel 是引用类型
	// 可以使用 for range 来迭代不断操作 channel
	// 可以设置单向或双向通道
	// 可以设置缓存大小，在未被填满前不会发生阻塞

	// 【Select】
	// 可处理一个或多个 channel 的发送与接收
	// 同时有多个可用的 channel时按随机顺序处理
	// 可用空的 select 来阻塞 main 函数
	// 可设置超时

	// 【方式1：time.Sleep】
	// go Go()
	// time.Sleep(2 * time.Second)

	// 【方式2：channel】
	// c := make(chan bool)
	// go func() {
	// 	fmt.Println("Go!")
	// 	c <- true
	// }()
	// // 阻塞，等待消息
	// <-c

	// 【方式3：】
	c := make(chan bool)
	go func() {
		fmt.Println("Go!")
		c <- true
		// 必须有 close，否则死锁
		close(c)
	}()
	// 不断存取
	for v := range c {
		fmt.Println(v)
	}
}

func Go() {
	fmt.Println("Go!")
}
