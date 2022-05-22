// 数据类型

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 一.整数
	// 类型   有无符号（有符号：前面可加负号 可表示负数） 占用存储空间(变量用此类型时所申请的内存空间大小) 可表示数值范围
	// uint8    无                                            1字节                           0到255
	// uint16   无                                            2字节                           0到65535
	// uint32   无                                            4字节                           0到4294967295
	// uint64   无                                            8字节                           0到18446744073709551615
	// int8     有                                            1字节                           -128到127
	// int16    有                                            2字节                           -32768到32767
	// int32    有                                            4字节                           -2147483648到2147483647
	// int64    有                                            8字节                           -9223372036854775808到9223372036854775807

	// int 有符号 32位操作系统和int32等价 64位操作系统和int64等价
	// uint 无符号 32位操作系统和uint32等价 64位操作系统和uint64等价
	// rune 有符号 与int32一样 表示一个Unicode码
	// byte 无符号 与uint8类似 存储字符时选用这个

	// 整数默认类型推导为int类型
	// var num1 = 10
	// 把num1的类型 填充到%T位置上
	// 计算出声明num1时所申请的内存大小 填充到%d上
	// 因为默认为int类型 所以num1变量占用内存空间为8个字节
	// fmt.Printf("num1类型为 %T 占用字节为 %d ", num1, unsafe.Sizeof(num1))

	var num2 int8 = 10
	// 因为声明变量时指定int8类型 所以num2申请的内存大小为1个字节
	// 所以在声明变量时 尽可能选用申请内存小的类型（当然也要考虑之后业务是否会加大此变量数值）
	fmt.Printf("num2类型: %T 占用字节: %d \n", num2, unsafe.Sizeof(num2))

	// 二.浮点数
	// 类型      有无符号   占用存储空间  表数范围
	// float32    有       4字节       -3.403E38到3.403E38
	// float64    有       8字节       -1.798E308到1.798E308

	var num3 float32 = 1.88
	fmt.Println("num3:", num3)

	num4 := 1.8888888888888
	// 类型自动推导为float64 小数位多的情况下float64比float32精度更准确
	// 日常开发中推荐使用float64位
	fmt.Printf("num4的类型: %T \n", num4)

	num5 := 5.12345e2
	// 相当于5.12345*10的2次方 e也可以大写 值是一样的
	fmt.Println("num5:", num5)

	num6 := 5.12345e-2
	// 相当于5.12345 / 10的2次方
	fmt.Println("num6:", num6)

	// 三.字符
	// golang中没有专门的字符类型 如果要保存单个字符（字母） 通常使用byte保存
	// 英文字母的字符 1字节  汉字的字符 3字节

	// 字符使用单引号包裹（字符本质是一个整数）
	// 变量存储的是字符在utf-8中的码值
	// 存储原理：字符->码值->二进制->存储
	// 读取原理：二进制->码值->字符->读取
	var char1 byte = 'a'
	// 直接输出byte类型变量 相当于输出的是变量值（字符）对应的utf-8码值
	fmt.Println("char1对应的码值:", char1)

	// 如果想直接输出字符 那么使用格式化输出
	fmt.Printf("char1的字符为: %c \n", char1)

	// var char2 byte = '呵'
	// 会报错 因为'呵'（21621）对应的utf-8码值 已经超过了byte类型最大值（255）
	// 在golang中字符编码都是使用utf-8来代表Unicode编码的

	// 使用int存储 因为int的存储范围 大于21621
	var char2 int = '呵'
	fmt.Println("char2的码值:", char2)

	// 如果存储的字符在ASCII表的 如：0-1 a-z A-Z 等 可直接使用byte 因为他们的码值 byte可以存放的下
	// utf-8和ASCII的关系是：utf-8是ASCII的升级版 utf-8中包含的字符更多（每个字符都有对应的utf-8码值） 同时utf-8中也包含ASCII中的字符
	// 如果码值较大 可以选择int等类型存储 utd-8字符码值查询网站：http://www.mytju.com/classcode/tools/encode_utf8.asp

	// 可以给变量赋值一个数字 按格式化（%c）输出 会输出该数字在utf-8中对应的字符
	var char3 int = 21621
	fmt.Printf("char3对应的字符为:%c \n", char3)

	// 字符是可以运算的
	var char4 = 10 + 'a'
	// a字符在utf-8中对应的码值是97 相当于 10+97=107
	fmt.Println("char4的码值:", char4)
	fmt.Printf("char4对应的字符为:%c \n", char4)

	// 四.bool
	// 布尔值占用一个字节
	// 布尔值只能是true或者false

	// 这样是错误的 不可改变数据类型
	// var isShow = true
	// isShow = 0

	// 五.字符串
	// 字符串就是一段固定长度的字符连接起来的字符序列

	// 这样是错误的
	// var str1 = "哈哈哈"
	// str1[0] = "嘿嘿"

	// 可以使用反引号输出源代码（以字符串的原生方式输出） 这样就不用各种转义了
	var str2 = `
    func main(){
			fmt.Println("main")
		}
	`
	fmt.Println("str2:", str2)

	// 字符串拼接
	var str3 = "哈哈哈" + "嘿嘿嘿"
	fmt.Println("str3:", str3)

	// 如果拼接需要换行的话 把加号留在上一行
	// 如果放在下一行开头 那么就会报错
	// 因为golang编译时会自动在语句末尾加分号
	// 如果看到加号表示这条语句还有下文 则不加分号 直到语句结束才会加
	// 但是 如果把加号放在下一行开头 那么第一行就会加分号 代表语句结束
	// 编译第二行代码的时候 开头多一个加号 就会报错
	// 例（错误的）：
	// var str4 = "哈哈哈"
	// + "嘿嘿嘿"
	// + "嘤嘤嘤"

	// 正确的：
	var str4 = "哈哈哈" +
		"嘿嘿嘿" +
		"嘤嘤嘤"
	fmt.Println("str4:", str4)
}
