// goPipeLine project main.go
package main

import (
	"fmt"
)

func main() {
	//	// Counter
	//	go func() {
	//		fmt.Println("naturals start")
	//		for x := 0; ; x++ {
	//			naturals <- x
	//		}
	//	}()
	//	// Squarer
	//	go func() {
	//		fmt.Println("Squarer start")
	//		for {
	//			x := <-naturals
	//			squares <- x * x
	//		}
	//	}()
	//	for {
	//		fmt.Println("main start")
	//		fmt.Println(<-squares)
	//		time.Sleep(time.Millisecond * 500)
	//	}
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
