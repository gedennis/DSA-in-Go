# 栈 stack

栈，是一种操作受限的线性表。它和队列的处理方式相反，栈是一种`后入先出(LIFO)`。

<img src="https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/06/03/screen-shot-2018-06-02-at-203523.png" alt="img" style="zoom:50%;" />

栈最常用的两个操作是`入栈`和`出栈`。在栈中插入数据叫做`入栈（push）`，总是在栈的末尾添加新元素。删除元素的操作叫做`出栈(pop)`，始终删除栈的最后一个元素（最新添加的元素）。

<img src="https://pic.leetcode-cn.com/691e2a8cca120acb18e77379c7cd7eec3835c8c102d1c699303f50accd1e09df-%E5%87%BA%E5%85%A5%E6%A0%88.gif" alt="img" style="zoom:50%;" />



## 实现

通常使用`动态数组`来实现栈结构。Go 语言中使用切片。





## 参考

- [Golang实现数据结构“栈”的三种实现，性能对比及应用示例](https://hansedong.github.io/2019/04/02/15/)
- [https://leetcode-cn.com/leetbook/read/queue-stack/ghqur/](https://leetcode-cn.com/leetbook/read/queue-stack/ghqur/)