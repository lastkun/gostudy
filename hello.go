package main

import "fmt"

// go语言中变量名在类型之前
func variable() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func main() {
	fmt.Println("hello world ")
	variable()
}
