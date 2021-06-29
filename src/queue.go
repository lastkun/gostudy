package main

import (
	"fmt"
	"lib/queue"
)

func main() {
	queues := queue.Queue{1, 2}
	fmt.Println(queues)
	queues.Push(3)
	fmt.Println(queues)
	queues.Pop()
	queues.Pop()
	queues.Pop()
	fmt.Println(queues.IsEmpty())
}
