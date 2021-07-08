package functional

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() intGen {
	a, b := 0, 1
	// 每次返回左边的数1,1,2,3,5,8
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	// 获取下一个元素
	next := g()
	//数字大于一万停止读取
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d \n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
