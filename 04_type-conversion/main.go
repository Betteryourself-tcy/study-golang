package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 一.数字类型相互转换
	// 编译错误 不同类型之间不能赋值
	// var num1 int32 = 100
	// var num2 float32 = i

	// 这样是正确的
	var num1 int32 = 100
	// 把num1的值转成float32类型 然后赋值给num2 不会影响num1本身的数据类型
	var num2 float32 = float32(num1)
	fmt.Printf("num1的值为:%v num1的类型为:%T num2的值为:%v num2的类型为:%T \n", num1, num1, num2, num2)

	// 转换时 必须是对方类型能接受的大小范围 编译虽然不会报错 但是结果是按溢出处理 和我们期待的结果不一样
	var num3 int64 = 10000000
	var num4 int8 = int8(num3)
	fmt.Println("num4:", num4)

	// 相加之后结果的类型（结果的类型为：相加的每个元素中 有明确定义类型的一方的类型）
	var num5 int8 = 10
	var num6 = num5 + 10
	fmt.Printf("num6的类型为：%T \n", num6)

	// 二.基本数据类型转为string
	// 方式一：fmt.Sprintf
	// 百分号解释：https://www.runoob.com/go/go-fmt-sprintf.html
	var num7 = 10
	var str1 = fmt.Sprintf("%d", num7)
	fmt.Printf("str1类型为:%T str1的值为%q \n", str1, str1)

	var num8 = 18.8
	var str2 = fmt.Sprintf("%f", num8)
	fmt.Printf("str2类型为:%T str2的值为%q \n", str2, str2)

	var bVal1 = true
	var str3 = fmt.Sprintf("%t", bVal1)
	fmt.Printf("str3类型为:%T str3的值为%q \n", str3, str3)

	var char = 'h'
	var str4 = fmt.Sprintf("%c", char)
	fmt.Printf("str4类型为:%T str4的值为%q \n", str4, str4)

	// 方式二：strconv
	// 详见文档 https://pkg.go.dev/strconv

	// 三.string转基本数据类型
	var str5 = "true"
	// golang中函数返回几个值就要接收几个值
	// 所以使用_忽略返回的第二个值 也就是不获取返回的第二个值
	var bVal2, _ = strconv.ParseBool(str5)
	fmt.Printf("bVal2的类型为:%T bVal2的值为:%v \n", bVal2, bVal2)

	var str6 = "100"
	// 10:十进制 64:int64位
	var num9, _ = strconv.ParseInt(str6, 10, 64)
	fmt.Printf("num9的类型为:%T num9的值为:%v \n", num9, num9)

	var str7 = "1.88"
	// 64:float64位
	var num10, _ = strconv.ParseFloat(str7, 64)
	fmt.Printf("num10的类型为:%T num10的值为:%v \n", num10, num10)

	// 要确保string类型能转成有效的数字
	// 比如可以把“123”可以转成123
	// 但是“哈哈哈”不能转成有效的数字 golang会把它转成int默认值0
	// 比如把“哈哈哈”转为布尔值 转换不成功 那么golang就会把它转成布尔值的默认值false
	var str8 = "哈哈哈"
	var num11, _ = strconv.ParseInt(str8, 10, 64)
	fmt.Printf("num11的类型为:%T num11的值为:%v \n", num11, num11)
}
