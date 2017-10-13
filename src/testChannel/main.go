// testChannel project main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var str string = "abcdefgh"
	start(str)
}

func start(str string) {
	size := make(chan rune)
	var wg sync.WaitGroup
	for _, f := range str {
		wg.Add(1)
		go func(r rune) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println(r)
			size <- r
		}(f)
	}

	go func() {
		wg.Wait()
		close(size)
	}()

	var total string
	fmt.Println("-=-=-=-=-")
	//	time.Sleep(time.Second * 3)
	//	close(size)
	for s := range size {
		total += string(s)
	}
	fmt.Println(total)
}
