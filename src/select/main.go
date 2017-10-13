// select project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		fmt.Printf("cur %d \n", i)
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
			fmt.Printf("xxx %d \n", i)
		default:
			fmt.Println("][")
		}

	}
}
