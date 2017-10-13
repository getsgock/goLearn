// chan2 project main.go
package main

import (
	"fmt"
	"sync"
	//	"time"
)

var c = make(chan int, 10)
var a string
var b string
var d = make(chan int)
var l sync.Mutex

func main() {
	//	l.Lock()
	//	go f()
	//	l.Lock()
	//	fmt.Println(a)
	//	twoprint()
	//	time.Sleep(time.Second)
	//	go setup()
	//	for g == nil {
	//	}
	//	print(g.msg)
	fmt.Println(test)
}

func f() {
	a = "=========="
	l.Unlock()
}

var once sync.Once
var done bool

type T struct {
	msg string
}

var g *T
var test = "123"

func setup() {
	//	done = true
	//	b = "hello, world-aaaaaaaaaaa\n"
	//	done = true
	//	fmt.Println(time.Now())
	t := new(T)
	t.msg = "hello, world"
	g = t
}
func doprint() {
	fmt.Println(done)
	if !done {
		once.Do(setup)
	}
	print(b)
}
func twoprint() {
	go doprint()
	go doprint()
}
