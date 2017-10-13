// face19 project main.go
package main

import (
	"fmt"
)

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
	s = false
	u
	t = iota
)

func main() {
	fmt.Println(x, y, z, k, p, s, u, t)
}
