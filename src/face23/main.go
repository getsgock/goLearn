// face23 project main.go
package main

func main() {
	for i := 0; i < 10; i++ {
	loop:
		println(i)
	}
	goto loop
}
