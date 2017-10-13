// unsafe project main.go
package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
	d float64
}

func main() {
	x.c = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 56, 7, 8, 9, 1}
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(x.a))
	fmt.Println(unsafe.Sizeof(x.b))
	fmt.Println(unsafe.Sizeof(x.c))
	fmt.Println(unsafe.Sizeof(x.d))
}
