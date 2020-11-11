package main

import "fmt"

type MyCircularQueue struct {
	data []int
	head int
	tail int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	var circularQueue = MyCircularQueue{
		data: make([]int, k),
		head: -1,
		tail: -1,
	}

	return circularQueue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	// 队满
	if this.IsFull() {
		return false
	}

	// 队空
	if this.IsEmpty() {
		this.head++
	}
	if this.tail == len(this.data)-1 {
		this.tail = 0
	} else {
		this.tail++
	}
	this.data[this.tail] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	// empty
	if this.IsEmpty() {
		return false
	}
	// last element
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}

	if this.head == len(this.data)-1 {
		this.head = 0
		return true
	}
	this.head++
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]

}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]

}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1 && this.tail == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail-this.head+1 == len(this.data)) || (this.head == this.tail+1)
}

func main() {
	var circularQueue = Constructor(3)
	cqPtr := &circularQueue
	fmt.Println(cqPtr.Front())
	fmt.Println(cqPtr.Rear())
	fmt.Println(cqPtr.EnQueue(1))
	fmt.Println(cqPtr.EnQueue(2))
	fmt.Println(cqPtr.EnQueue(3))
	fmt.Println(cqPtr.EnQueue(4))
	fmt.Println(cqPtr.Rear())
	fmt.Println(cqPtr.IsFull())
	fmt.Println(cqPtr.DeQueue())
	fmt.Println(cqPtr.EnQueue(4))
	fmt.Println(cqPtr.Rear())
	fmt.Println(cqPtr.EnQueue(4))
}
