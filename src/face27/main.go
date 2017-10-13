// face27 project main.go
package main

import (
	"errors"
	"fmt"
)

var ErrDidNotWork = errors.New("did not work")

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}
func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}
