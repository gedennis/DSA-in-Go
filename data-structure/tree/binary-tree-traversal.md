# 二叉树的遍历

前面讲了二叉树的一些基本概念，这一次我们来讲讲二叉树的遍历。二叉树的遍历很重要，它的基本思想和算法套路基本可以用在所有二叉树相关的问题。

## 遍历方法

遍历简单来说，就是把树的节点都走一遍。常用的二叉树遍历有三种：**前序遍历**、**中序遍历**和**后序遍历**。

- 前序遍历：对于任意节点，先该节点，然后左子树，最后右子树

- 中序遍历：左子树，节点，右子树

- 后序遍历：左子树，右子树，节点。

下面这个图会更清楚一点，大家记得一定要自己去走一遍：

<img src="https://static001.geekbang.org/resource/image/ab/16/ab103822e75b5b15c615b68560cb2416.jpg" alt="img" style="zoom:50%;" />

> 还有一种遍历方式：层序遍历。层序遍历是从根节点开始，按层进行遍历的，通常以队列实现，这里先不做介绍。

## Go 语言实现

二叉树的前、中、后序遍历就是一个递归的过程。只要掌握根节点的遍历方式，后面就用递归的方式来遍历左右子树就好了

### 基本结构

在开始遍历之前，我们先简单用 go 代码来生成一棵树。这里我们链表来生成树，每个节点的节点如下，以 int 型数据为例：

```go
type TreeNode struct {
	Val   int           // 值
	Left  *TreeNode     // 左子树指针
	Right *TreeNode     // 右子树指针
}
```

生成一个这样的树：

<img src="https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg" alt="img" style="zoom: 50%;" />

示例代码如下，大家可以根据自己要实现的树来实现：

```go
var root *TreeNode	
root = &TreeNode{Val: 1}
root.Right = &TreeNode{Val: 2}
root.Right.Left = &TreeNode{Val: 3}
```

基础代码如下：

```go
package main
import "fmt"

// 树的基本结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 生成几种不同类型的树，大家可以自己实现
func makeTree(kind int) *TreeNode {
	var root *TreeNode

	switch kind {
	case 1:
		root = &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Left = &TreeNode{Val: 3}
	case 2:  // 只有根节点
		root = &TreeNode{Val: 1}
	default: // nil 树

	}
	return root
}
func main() {
  // 生成树，第一种类型的树
	root := makeTree(1)
  // TODO: 在这里做遍历
}
```

这里我们使用一个 `makeTree` 函数，根据 kind 来实现不同的树。

### 前序遍历

下面以 leetcode 144题 为例来实现，大家可以[查看这里](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)。

```go
func preorderTraversal(root *TreeNode) []int {
  // 递归终止
	if root == nil {
		return []int{}
	}
  // 根节点
	result := []int{root.Val}
	// 左子树
	left := preorderTraversal(root.Left)
	result = append(result, left...)
	// 右子树
	right := preorderTraversal(root.Right)
	result = append(result, right...)

	return result
}

```

### 中序遍历

以 leetcode 94题为例，[链接地址](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)。

```go
func inorderTraversal(root *TreeNode) []int {
  // 递归终止
	if root == nil {
		return []int{}
	}
	result := []int{}
	// 1. 左子树
	left := inorderTraversal(root.Left)
	result = append(result, left...)
	// 2. 根节点
	result = append(result, root.Val)
	// 3. 右子树
	right := inorderTraversal(root.Right)
	result = append(result, right...)

	return result
}
```

### 后序遍历

以 leetcode 144题为例，[链接地址](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)

```go
func postorderTraversal(root *TreeNode) []int {
  // 递归终止
	if root == nil {
		return []int{}
	}
	result := []int{}
	// 1. 左子树
	left := postorderTraversal(root.Left)
	result = append(result, left...)
	// 2. 右子树
	right := postorderTraversal(root.Right)
	result = append(result, right...)
	// 3. 根节点
	result = append(result, root.Val)

	return result
}
```



### 规律

大家写到这里有没有发现，二叉树算法的递归实现都有一个套路。

```go
func Operation(root *TreeNode) {
  // 递归终止
  if (root == nil) {
    // 相应的返回
  }
  // 根节点的处理（前序）
  
  // 左子树处理
  Operation(root.Left)
  
  // 根节点的处理（中序）
  
   // 右子树处理
  Operation(root.Right)
  
   // 根节点的处理（后序）
  
   // 整理结果返回
}
```

如上，大部分结构是固定的，只有根节点的处理，根据不同的需求，采取前序、中序或者后序三种遍历方式的一种来实现。后面我们相关有二叉树的题基本都是这个思路。





