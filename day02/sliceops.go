package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 10, 20)
	printSlice(s1)
	printSlice(s2)

	// copy 操作
	copy(s2, s1)

	// 删掉某个元素  s2[4:]... 代表切片内的所有元素
	s2 = append(s2[:3], s2[4:]...)

	printSlice(s2)
}

func printSlice(slice []int) {
	fmt.Printf("slice = %v , len = %d cap = %d", slice, len(slice), cap(slice))
	fmt.Println()
}
