// 变量

package main

import "fmt"

// 一.全局变量
// 这是全局变量 类型可不写 根据值自动推断变量类型（类型推导）
var testOne string = "哈哈哈"

// 二.默认值
// 如果变量没有赋值
// int float类型默认值为0
// string类型默认值为“”
// bool类型默认为false
var num int

// 三.多个变量声明方式
// 类型相同
// var num1, num2 int
// 类型推导
// var num1, str = 1, "better-yourself"
var (
	num1 = 1
	str  = "better-yourself"
)

// 四.全局常量
const testTwo = "嘿嘿"

func knowledge() {
	// 五.局部变量
	// 局部变量可以使用 name := “哈哈哈” 再次赋值直接使用name="嘿嘿嘿"就可
	name := "哈哈哈"
	// name = 123 这是错误的 不可更改数据类型
	name = "呵呵呵"
	fmt.Println(name)
}

func main() {
	fmt.Println(testOne, testTwo, num, num1, str)
	knowledge()
}
