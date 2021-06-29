package main

import "fmt"

// 函数外面必须用var 等关键字定义
var (
	i = 3
	j = "s"
)

// go语言中变量名在类型之前
func variable() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 第一次定义变量要用 :=
func variableDiverType() {
	a, b, c, d := 3, 4, true, "s"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("hello world ")
	variable()
	variableDiverType()
	fmt.Println(i, j)
}
