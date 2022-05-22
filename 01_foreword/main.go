// 前言

// 官网下载go版本 1.14以后不用配置环境变量 然后在vscode中安装Go插件 即可编写go程序
// 使用go module构建项目：go mod init xxx
// go中一个文件夹代表一个包 一个包一级目录下package定义的源文件名（包名）及函数名不可重复
// 但是go中可以存在子包（子文件夹）子包中package定义的源文件名（包名）及函数名和父文件夹中定义的并不冲突
// go中包名不可以有中文
// 每个项目一般主文件命名为main.go 定义源文件名（包名）为package main 里面包含一个main函数（主函数）
// 在go文件中输入快捷命令 pkgm 可快速生成一个基础代码体
// 打印输出快捷键：fp ->	fmt.Println("哈哈哈")
// 格式化输出快捷键：ff -> fmt.Printf("num1类型为 %T 占用字节为 %d ", num1, unsafe.Sizeof(num1))
// unsafe.Sizeof(num1) 计算变量占用的字节（声明变量时 申请的内存空间大小）
// 上述Printf函数的第二个参数会填充到%T位置上 第三个参数会填充到%d位置上
// 1.%T -> 变量类型 2.%d -> 变量申请内存大小 3.%v -> 变量值
// command/ctrl + s 保存和自动格式化代码 代码中使用的包 会自动导入
// fmt是go系统中自带的一个包
// go中不允许使用单引号（除单个字符外）
// go中引入的包或声明的变量未使用的情况下 会报错
// vscode安装Code Runner 点击右上角类似于开始的按钮 即可运行程序 等同于 go run main.go
// 在终端中跑go项目 使用go run main.go（本质上run内部也是先把.go文件转成二进制文件 然后去执行的）
// 打包用go build xxx.go生成二进制可执行文件 也可指定生成文件名 go build -o 自定义文件名.exe xxx.go
// windows终端下执行二进制文件 xxx.exe
// mac linux终端下执行二进制文件 ./xxx

// dos常用指令
// 一.目录操作
//   1.查看当前目录及当前目录包含的文件有哪些：dir
//   2.切换到F盘：cd /d f:
//   3.回到上级目录：cd ..
//   4.回到根目录（f盘 d盘等等）：cd \
// 二.文件夹操作
//   1.创建目录：md xxx
//   2.删除空目录：rd xxx
//   3.删除目录及目录内所有：rd /q/s xxx (/q:不带询问 /s:递归删除所有)
// 三.文件操作
//   1.往文件中写入内容（如果没有此文件，系统则默认创建）：echo better-yourself > xxx.txt
//   2.复制文件到xxx路径：copy xxx.txt 文件夹路径\ok.txt
//   (\ok.txt:复制过来的文件重命名为ok.txt,不写则使用原文件名)
//   3.剪切文件到xxx路径P：move xxx.txt 文件夹路径\ok.txt(\ok.txt同上)
//   4.删除某个文件：del xxx.txt(输入文件名时 按tab键 名称可自动补全)
//   5.删除所有txt结尾的文件：del *.txt
// 四.其他指令
//   1.清屏：cls
//   2.退出dos：exit

package main

func main() {

}
