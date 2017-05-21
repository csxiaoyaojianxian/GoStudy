/*
* @Author: sunshine
* @Date:   2017-04-29 22:44:57
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-29 23:50:08
 */

package main

import (
	"fmt"
	"os"
)

func cmdHelp() {
	fmt.Print("***********************Help**********************\n")
	fmt.Print("*\t\t\t\t\t\t*\n")
	fmt.Print("*\thelp\tShow help list\t\t\t*\n")
	fmt.Print("*\tls\tList files\t\t\t*\n")
	fmt.Print("*\tpwd\tPrint working directory\t\t*\n")
	fmt.Print("*\tmkdir\tCreate directory\t\t*\n")
	fmt.Print("*\tcp\tCopy file\t\t\t*\n")
	fmt.Print("*\tmv\tMove or rename\t\t\t*\n")
	fmt.Print("*\trm\tDelete file\t\t\t*\n")
	fmt.Print("*\tprint\tPrint Code\t\t\t*\n")
	fmt.Print("*\tquit\tQuit menu program\t\t*\n")
	fmt.Print("*\t\t\t\t\t\t*\n")
	fmt.Print("*************************************************\n\n")
}
func system(cmd string) {
	switch cmd {
	case "ls":
		fmt.Println("system.ls")
	case "pwd":
		fmt.Println("system.pwd")
	case "mkdir":
		fmt.Println("system.mkdir")
	case "cp":
		fmt.Println("system.cp")
	case "mv":
		fmt.Println("system.mv")
	case "rm":
		fmt.Println("system.rm")
	default:
		fmt.Println("undefined")
	}
}

func printCode() {
	userFile := "lab2.go"
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func main() {
	var cmd string
	cmdHelp()
	for {
		fmt.Print("Menu->")
		fmt.Scanf("%s\n", &cmd)
		if cmd == "help" {
			cmdHelp()
		} else if cmd == "ls" {
			system(cmd)
		} else if cmd == "pwd" {
			system(cmd)
		} else if cmd == "mkdir" {
			system(cmd)
		} else if cmd == "cp" {
			system(cmd)
		} else if cmd == "mv" {
			system(cmd)
		} else if cmd == "rm" {
			system(cmd)
		} else if cmd == "print" {
			printCode()
		} else if cmd == "quit" {
			break
		} else {
			fmt.Print("Error: unsupported command, you can use 'help' to list the available commands\n")
		}
	}
}
