package main

import (
	"fmt"
	"time"
)

// chan<- 表示该函数返回的channel是send-only的 使用者只能发送数据给这个channel，不能从里面获得数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()

	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	// 给channel输送数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //这里是int类型的接收者，'a'+i转为了ascii码存储
	}
	// sleep的原因是 main协程执行结束后会销毁所有子协程
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
