// face15 project main.go
package main

import (
	"fmt"
)

func main() {
	list := new([]int)
	list = append(list, 1) //first argument to append must be slice; have *[]int
	fmt.Println(list)
	//切片需要make
}
