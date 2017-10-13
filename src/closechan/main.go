// closechan project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan bool)
	close(c)
	for {
		time.Sleep(time.Second * 1)
		select {
		case x := <-c:
			fmt.Println(x)
		default:
			fmt.Println("-=-=")
		}
	}
}
