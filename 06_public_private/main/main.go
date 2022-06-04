// 公有变量和私有变量介绍

package main

import (
	"fmt"
	"utils"
)

func main() {
	// name not exported by package
	// fmt.Println(utils.name)

	fmt.Println(utils.Name)
}
