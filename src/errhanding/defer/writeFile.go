package writeFile

import (
	"bufio"
	"fmt"
	"lib/functional"
	"os"
)

func WriteFile(filename string) {
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	//最后将buffer里的内容写到文件，在file.Close()之前执行
	defer writer.Flush()

	fibonacci := functional.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fibonacci())
	}
}
