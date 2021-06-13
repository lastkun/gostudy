package main

import (
	"fmt"
	"math"
)

func test() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enums() {
	const (
		a = iota // iota ：按规则递增 从0开始
		_
		c
		d
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)

	fmt.Println(a, c, d)
	fmt.Println(b, kb, mb, gb, tb)
}

func consts() {
	const (
		name = "abc"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(name, c)
	enums()
}

func main() {
	test()
	consts()
}
