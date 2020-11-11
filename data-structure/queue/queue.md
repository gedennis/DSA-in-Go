# 队列

在数组中，可以根据索引随机访问元素。但在某些场景下，我们需要限制处理元素的顺序。两种典型的操作受限的线性表结构是`队列`和`栈`。本章将主要介绍队列的相关概念。

## 什么是队列？

`队列(Queue)`是一种`先进先出 FIFO`的数据结构，首先处理添加到队列中的第一个元素。

<img src="https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/08/14/screen-shot-2018-05-03-at-151021.png" alt="img" style="zoom:50%;" />

插入操作叫做`入队（enqueue）`,新元素添加到队列的末尾。删除操作叫做`出队（dequeue)`，只能移除第一个元素。

![queue](https://pic.leetcode-cn.com/44b3a817f0880f168de9574075b61bd204fdc77748d4e04448603d6956c6428a-%E5%87%BA%E5%85%A5%E9%98%9F.gif)

## 实现

为了实现队列，我们使用`动态数组`（Go语言中为切片）和一个指向队列头部的索引。队列支持两种操作：`入队`和出队。

```go
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
```

缺点：随着起始指针的移动，浪费了越来越多的存储空间。

## 循环队列

上面实现的队列比较低效，更有效的办法是使用`循环队列`。方法：

- 使用固定大小的数组存储数据
- 使用两个指针表示起始位置 head 和 结束位置 tail

这样可以解决上面实现存储空间浪费的问题。

**实现**

```go
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
```





## 参考

- [The complete guide to Go Data Structures](https://flaviocopes.com/golang-data-structures/)
- 