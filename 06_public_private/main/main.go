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
