// 值类型地址值 指针 值类型和引用类型的介绍

package main

import "fmt"

func main() {
	// 一.获取666内存空间对应的地址值
	var num int = 666
	// 0x...
	fmt.Println("num的地址值:", &num)

	// 二.指针变量
	// ptr是一个指针变量
	// 类型为 *int
	// 保存的值是一个地址值
	var ptr *int = &num
	fmt.Println("ptr的值是:", ptr)
	// ptr保存一个地址值 同样也是需要开辟一块内存空间的
	// 0x...
	fmt.Println("ptr的地址值是:", &ptr)

	// *ptr 获取ptr保存的地址值所对应的具体值（对应的那块内存空间保存的值）
	// ptr 获取到地址值
	// * 通过地址值 获取到对应内存空间中的具体值
	// 打印值为666
	fmt.Println("ptr保存的地址值所指向的真实值:", *ptr)

	// 三.通过地址值 改变内存空间中的具体值
	var num2 = 10
	// 获取地址值
	var ptr2 = &num2
	// 通过地址值访问内存空间 并修改
	*ptr2 = 100
	fmt.Println("num2:", num2)

	// 四.其他值类型对应的地址值
	var str = "哈哈哈"
	var ptr3 = &str
	// 0x...
	fmt.Println("ptr3的值是:", ptr3)

  // 五.值类型和引用类型的介绍
	// 值类型包括：int float string bool及数组和结构体
	// 引用类型包括：指针 切片 map chan 接口 等

	// 值类型：变量直接存储值 内存通常在栈中分配
	// 引用类型：变量存储的是一个地址值 这个地址值对应的内存空间存储真正的值
	// 地址值对应的内存空间通常在堆上进行分配
	// 当没有任何变量引用这个地址值时 该地址值对应的内存空间 就会把GC（垃圾回收器）回收
}