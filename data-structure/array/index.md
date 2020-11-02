# 数组 Array

`数组（Array）`是一种**线性表**数据结构。它用一块`连续的内存空间`，来存储一组具有`相同数据类型`的数据。

> 注意：
>
> - 这里的数组是指一种**数据结构**，和编程语言中的**数据类型**数组是两个概念，不能混淆。
> - 数据结构的角度，顺序表这个概念会更清楚一点。所以本文约定使用顺序表这个概念。
> - 顺序表是通过数组来编程实现的。

例如，一个长度为10 的 int 类型的顺序表，每个元素占 4 个字节，分配连续内存空间 1000 - 1039。

<img src="https://static001.geekbang.org/resource/image/98/c4/98df8e702b14096e7ee4a5141260cdc4.jpg" alt="img" style="zoom:50%;" />

## 常见操作

### 读取元素

因为数组元素存储在连续的内存空间，以及具有相同数据类型，所以可以快速的得到顺序表中任一元素的内存地址。所以顺序表支持`随机访问`，通过索引快速访问，时间复杂度为常数级 O(1)。

```go
a[i]_address = a[0]_address + i * 数据类型的大小
```

以数组 `["C", "O", "D", "E", "R"]` 为例，它的各元素对应的索引及内存地址如下图所示。

<img src="https://pic.leetcode-cn.com/273ac74bdd7a19d72c2bf60d84ddd66f09b45de4d8c36333bf5f1fee2c7a8330-%E5%9B%BE%E7%89%872.png" alt="2.png" style="zoom:50%;" />

假如我们想要访问索引为 2 处的元素 "D" 时，计算机会进行以下计算：

- 找到该数组的索引 0 的内存地址： 2008；
- 将内存地址加上索引值，作为目标元素的地址，即 2008 + 2 = 2010，对应的元素为 "D"，这时便找到了目标元素。

### 查找元素

如果要查找数组的某个元素，我们需要从数组开头逐步向后查找。如果找到目标元素，就停止查找。

![](https://pic.leetcode-cn.com/3d9c20552e0e9c4650f4a267f4066aa71338ad0013514559b57a1bf786d662ba-4.gif)

当目标元素为第一个元素时，时间复杂度最优为 O(1)。当目标元素为最后一个或者不存在时，时间复杂度为 O(n)。平均复杂度为 O(n)。

### 插入元素

假如 数组a 长度为n。如果需要在位置 K 插入元素，需要将 k 到 n 的元素后移，然后插入新元素。时间复杂度为 O(n)。如果是最后一个元素，最优，时间复杂度为O(1)。

![7.gif](https://pic.leetcode-cn.com/22ce7dbf8cd441fd7425499cd8154d1c4211a6a42ec3f3995520ee76ce7183c7-7.gif)

如果需要对数组进行频繁的插入操作，会造成时间的浪费。

优化：对于无序数组。

### 删除元素

和插入操作类似，如果要删除位置为K的元素，需要把 K+1 到 n 的元素前移。以删除索引 `1` 中的元素 `"O"` 为例，具体过程如图所示。

![6.gif](https://pic.leetcode-cn.com/4df7a5a75e5f76b6e7e4540f9403c7c2fee5197a1f30421b4f5d32fdca2cf360-8.gif)

最后一个元素最优，时间复杂度为O(1)。平均复杂度 O(n)。

## Go 语言实现

### Go 数组与切片

Go 语言提供数组这种数据类型。不过数组的大小是固定的，不支持动态扩容。

Go 提供另外一种数据类型`切片（Slice）` ，它相当于动态数组，可以动态扩容。当数据满了之后再插入元素，会重新申请一块 2 倍的内存空间，然后将原数组元素复制过去。

## 参考

【1】[https://leetcode-cn.com/leetbook/read/array-and-string/yjcir/](https://leetcode-cn.com/leetbook/read/array-and-string/yjcir/)  
【2】[数据结构与算法之美 - 数组](https://time.geekbang.org/column/article/40961)