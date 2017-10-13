// face12 project main.go
package main

func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
func GetValue() int {
	return 1
}
