package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 十进制转二进制输出
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}

	scanner := bufio.NewScanner(file)

	// 如果只有终止条件 可以写两个分号
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	fmt.Println(convertToBin(13))
	printFile("a.txt")
}
