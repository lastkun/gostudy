package main

import "fmt"

func main() {
	// 定义map
	var map1 map[string]int      // nil
	map2 := make(map[string]int) // empty map
	// 每一行后都带逗号
	map3 := map[string]string{
		"name": "kun",
		"age":  "18",
	}

	fmt.Println(map1, map2, map3)
	//遍历
	for k, v := range map3 {
		fmt.Println(k, v)
	}
	// 取值
	fmt.Println(map3["name"])
	fmt.Println(map3["hehe"]) // 输出一个空串(zero value)

	//判断是否存在key
	value, flag := map3["name"]

	fmt.Println(value, flag)

	if value, flag = map3["agee"]; flag {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exit")
	}

	// 删除
	delete(map3, "age")
	fmt.Println(map3)

}
