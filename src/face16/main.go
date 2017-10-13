// face16 project main.go
package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2) //cannot use s2[0:] (type []int) as type int in append
	fmt.Println(s1)
}
