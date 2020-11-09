

递归是解决树的相关问题最有效和最常用的方法之一。通常可以通过两种方式的递归来解决问题：

- 自顶向下
- 自底向上

## 自顶向下

自顶向下，是指每次递归，先对节点做处理得到某些值，然后将这些值当做参数，传给子节点的递归。这个过程类似于`前序遍历`。递归函数`top_down(root, params)`具体框架如下：

```
1. return specific value for null node
2. update the answer if needed                      // anwer <-- params
3. left_ans = top_down(root.left, left_params)		// left_params <-- root.val, params
4. right_ans = top_down(root.right, right_params)	// right_params <-- root.val, params
5. return the answer if needed                      // answer <-- left_ans, right_ans
```

### 例子: 二叉树的最大深度

以求为例。我们知道根节点的深度为 1。对于任意节点，如果知道该节点的深度 为 k，那它的子节点的深度 为 K + 1。所以，我们可以在调用递归函数的时候，将节点的深度作为一个参数传递。直到叶子节点，我们可以比较当前叶子节点和当前最大深度，更新答案。伪代码如下：

```go
1. return if root is null
2. if root is a leaf node:
3. 		answer = max(answer, depth)         // update the answer if needed
4. maximum_depth(root.left, depth + 1)      // call the function recursively for left child
5. maximum_depth(root.right, depth + 1)		// call the function recursively for right child
```

 ### Go 实现

```go
func maxDepth(root *TreeNode) int {
	depth := 0
	var findDepth func(base *TreeNode, currentDepth int)
	findDepth = func(base *TreeNode, currentDepth int) {
		if base == nil {
			return
		}
		// leaf node
		if base.Left == nil && base.Right == nil {
			if currentDepth > depth {
				depth = currentDepth
			}
		}

		findDepth(base.Left, currentDepth+1)
		findDepth(base.Right, currentDepth+1)
	}
	findDepth(root, depth+1)
	return depth
}
```



## 自底向上

自底向下是指，每轮递归，先对子节点递归，找到子树的求解。最后根据当前节点的值和子树的解，综合得到问题的解。这个过程类似于`后序遍历`。“自底向上” 的递归函数 `bottom_up(root)` 为如下所示：

```
1. return specific value for null node
2. left_ans = bottom_up(root.left)			// call function recursively for left child
3. right_ans = bottom_up(root.right)		// call function recursively for right child
4. return answers                           // answer <-- left_ans, right_ans, root.val
```

### 例子: 二叉树的最大深度

我们知道根节点的最大深度，等于左子树的最大深度和右子树的最大深度的最大值再加上 1。所以，如果我们通过递归知道左子树和右子树的深度就可以了，一次类推。

下面是递归函数 `maximum_depth(root)` 的伪代码：

```go
1. return 0 if root is null                 // return 0 for null node
2. left_depth = maximum_depth(root.left)
3. right_depth = maximum_depth(root.right)
4. return max(left_depth, right_depth) + 1	// return depth of the subtree rooted at root
```

### go 实现

```go
// max depth, from down to top
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth2(root.Left)
	rightDepth := maxDepth2(root.Right)

	depth := leftDepth
	if rightDepth > leftDepth {
		depth = rightDepth
	}

	return depth + 1
}
```



## 两种方法对比

我个人会更喜欢更喜欢自底向上的方法，这种方法代码更简洁。那何时使用自底向上呢？

**自底向上**：对于树的任意节点，如果知道子树的解，是不是可以通过某个逻辑得到该节点的解。如果可以，则可以利用自底向上的方法求解。

**自顶向下**：能通过节点本身来得到某个子解，然后通过确定某些参数，传递给子树的递归，直到找到整棵树问题的解。



## 参考

【1】[https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xefb4e/](https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xefb4e/)