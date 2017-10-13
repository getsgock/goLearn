// defer project main.go
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	bigSlowOperation()
}
func bigSlowOperation() {
	//	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	defer trace4()()
	fmt.Println("--------------")
	time.Sleep(3 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func trace2() {
	fmt.Println("!!!!!!!!!!!!!!")
}
func trace3() func() {
	fmt.Println("!!!!!!!!!!!!!!")
	return func() func() {
		fmt.Println("@@@@@@@@@@@@@@")
		return func() {
			fmt.Println("##############")
		}
	}()
}

func trace4() func() {
	defer func() {
		fmt.Println("22222222222")
	}()
	fmt.Println("1111111111111")
	return func() {
		fmt.Println("33333333333")
	}
}

func daax() (eax int, err error) {
	eax = 3
	err = fmt.Errorf("111111111111111")
	return eax, err
}
