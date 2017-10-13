// face29 project main.go
package main

import (
	_ "fmt"
)

func main() {
	a, b := test(100)
	a()
	b()
}
func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}
