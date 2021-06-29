package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func eval(a, b int, o string) (int, error) {
	switch o {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf(
			"error operation : %s", o)
	}
}

//函数式编程
func apply(op func(int, int) int, a, b int) int {
	//用反射拿函数名
	pointer := reflect.ValueOf(op).Pointer()
	funName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function: %s", funName)
	return op(a, b)
}

//可变参数列表
func sum(numbers ...int) int {
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	q, r := div(10, 3)
	fmt.Println(q, r)

	if result, e := eval(3, 4, "x"); e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4))

}
