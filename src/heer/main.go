// heer project main.go
package main

import (
	"fmt"
	//	"reflect"
	//	aa "unsafe"
)

type book struct {
	a bool
	b int
}

func main() {
	a := byte(3)
	b := byte(5)
	fmt.Printf("a -> %08b \n", a)
	fmt.Printf("b -> %08b \n", b)
	fmt.Printf("a &^ b -> %08b \n", (a &^ b))
	fmt.Printf("b &^ a -> %08b \n", (b &^ a))
	fmt.Printf("a ^ b ->% 08b \n", (b ^ a))
	fmt.Printf("a & b ->% 08b \n", (b & a))
	fmt.Printf("a | b ->% 08b \n", (b | a))
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	fmt.Printf("%b  \n", x>>(0*8))
	fmt.Printf("%b  \n", byte(x>>(0*8)))
	fmt.Println("n is ", x)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
