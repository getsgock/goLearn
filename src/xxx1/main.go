// xxx1 project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		httpHandler()
	}

}

func httpHandler() {
	errCh := make(chan error, 1)
	resCh := make(chan int, 1)
	go func(ch1 chan error, ch2 chan int) {
		defer close(ch1)
		defer close(ch2)
		ch1 <- fmt.Errorf("shit")
	}(errCh, resCh)
	time.Sleep(time.Millisecond * 50)
	select {
	case <-errCh:
	case <-resCh:
		fmt.Println("2141231231")
	}
}
