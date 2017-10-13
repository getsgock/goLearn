// mutex project main.go
package main

import (
	"fmt"
)

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func main() {
	fmt.Println("Hello World!")
}

func 