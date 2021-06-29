package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}

	updateArr(arr[:])
	fmt.Println(arr)

	//切片还可以再切片

	s := arr[2:]
	s = s[:5]

	fmt.Println(s)

	//
	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Printf("s1 = %v , len = %d , cap = %d", s1, len(s1), cap(s1))
	fmt.Println()
	fmt.Printf("s2 = %v , len = %d , cap = %d", s2, len(s2), cap(s2))
}

//这里的参数是切片形式
func updateArr(arr []int) {
	arr[0] = 100
}
