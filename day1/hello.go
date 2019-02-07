// go 语言以包作为管理单位
// 每个包必须声明包
// 每个包必须有一个main包
package main

import "fmt"

func main() { // 左括号必须与函数名同行
	// 调用函数，大部分都需要导入包
	/* 完美的注释
	 */
	fmt.Printf("hello, world\n")
	fmt.Printf("hello")
}
