package queue

import "fmt"

type Queue []int

// 指针
func (q *Queue) Push(value int) {
	fmt.Println(q)
	fmt.Println(&q)
	*q = append(*q, value)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
