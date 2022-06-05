// 终端交互
// 以下代码需要在08文件夹终端中 使用go run main.go 来运行
package main

import "fmt"

func main() {
	var tool string

	fmt.Println("请输入vite或webpack")

	// 获取终端中输入的内容 赋值给tool变量
	// 代码运行到这行的时候
	// 程序会等待用户输入并按回车 才会继续向下执行
	fmt.Scanln(&tool)

	fmt.Println("选择的工具是:", tool)

}
