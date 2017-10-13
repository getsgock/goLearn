// face25 project main.go
package main

import (
	"fmt"
)
type User struct {
}
type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
    fmt.Println("MyUser1.m1")
}
func (i User) m2(){
    fmt.Println("User.m2")
}
//考点：**Go 1.9 新特性 Type Alias **因为MyUser2完全等价于User，
//所以具有其所有的方法，并且其中一个新增了方法，另外一个也会有。 但是
func main() {
	var i1 MyUser1
    var i2 MyUser2
    i1.m1()
    i2.m2()
	//MyUser1.m1  User.m2

}
