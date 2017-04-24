/*
* @Author: sunshine
* @Date:   2017-04-22 22:19:11
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-23 14:39:53
 */

package main

import "fmt"

func main() {
	/*
	   【Go环境变量与工作目录】
	   GOPATH下需要建立3个目录：
	     bin（存放编译后生成的可执行文件）
	     pkg（存放编译后生成的包文件）
	     src（存放项目源码）

	   【Go命令】
	   Go常用命令简介
	     go get：获取远程包（需 提前安装 git或hg）
	     go run：直接运行程序
	     go build：测试编译，检查是否有编译错误
	     go fmt：格式化源码（部分IDE在保存时自动调用）
	     go install：编译包文件并编译整个程序
	     go test：运行测试文件
	     go doc：查看文档（CHM手册）

	   【sublime】
	   安装 gosublime

	   【调试】
	   sublime 中 ctrl+B

	*/

	fmt.Println("Hello World!")
}
