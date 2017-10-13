// GOMAXPROCS project main.go
package main

import (
	"fmt"
	//	"runtime"
)

func main() {
	//	runtime.GOMAXPROCS(3)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
