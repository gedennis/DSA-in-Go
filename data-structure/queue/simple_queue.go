package main

import "fmt"

// 定义队列
type MyQueue struct {
	data   []int
	pStart int
}

func (q *MyQueue) isEmpty() bool {
	return q.pStart >= len(q.data)

}
func (q *MyQueue) enQueue(val int) bool {
	// 未初始化
	if q.data == nil {
		q.data = []int{}
	}
	q.data = append(q.data, val)
	return true
}
func (q *MyQueue) deQueue() bool {
	if q.isEmpty() {
		return false
	}
	q.pStart++
	return true
}
func (q *MyQueue) getFront() int {
	return q.data[q.pStart]
}
func main() {
	var q MyQueue
	q.enQueue(5)
	q.enQueue(3)
	if !q.isEmpty() {
		fmt.Println("队首元素：", q.getFront())
	}
	q.deQueue()
	if !q.isEmpty() {
		fmt.Println("队首元素：", q.getFront())
	}
	q.deQueue()
	if !q.isEmpty() {
		fmt.Println("队首元素：", q.getFront())
	}
}
