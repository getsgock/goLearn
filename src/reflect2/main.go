// reflect2 project main.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()
	fmt.Println(x)
	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())
	d.Set(reflect.ValueOf(3))
	fmt.Println(d)
}
