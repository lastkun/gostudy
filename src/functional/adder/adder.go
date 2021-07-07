package main

import "fmt"

func adder() func(int) int {
	//sum 自由变量
	sum := 0
	// v 局部变量
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(v int) iAdder {
	return func(base int) (int, iAdder) {
		return v + base, adder2(base + v)
	}
}

func main() {
	adder1 := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(adder1(i))
	}
}
