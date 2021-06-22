package main

import "fmt"

func main() {
	//定义数组 数量要写在类型前面
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8}

	fmt.Println(arr1, arr2, arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 数组的遍历 range关键字  i : 下表 v ： 值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 记录最大值及下标
	var maxIndex int
	var maxValue int

	for i, v := range arr3 {

		if v > maxValue {
			maxIndex, maxValue = i, v
		}
	}
	fmt.Println(maxIndex, maxValue)

	//累加
	var sum = 0

	for _, v := range arr3 {
		sum += v
	}

	fmt.Println(sum)
}
