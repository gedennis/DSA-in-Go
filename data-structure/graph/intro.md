# 图 Graph

 图，是一种非线性数据结构。如下图所示：

<img src="https://static001.geekbang.org/resource/image/df/af/df85dc345a9726cab0338e68982fd1af.jpg" alt="img" style="zoom:50%;" />

图中的元素叫做`顶点(Vertext)`。顶点可以与其他任意节点建立联系（这一点和树不同），建立的联系叫做`边（Edge）`。顶点跟其他节点的边数，叫做`度(Degree)`。

图的一个典型应用就是社交网络，每个人是一个顶点，可以与其他人（顶点）互加好友。每添加一个好友就建立一条边。好友的个数就是`顶点的度`。大家还比较熟悉关注的功能，但是有单向关注和互关，怎么办呢？这里我们就要引入另一个概念：`有向图`。

<img src="https://static001.geekbang.org/resource/image/c3/96/c31759a37d8a8719841f347bd479b796.jpg" alt="img" style="zoom:50%;" />

如果边有方向，则为有向图，否则为无向图。有向图中，出去的边叫做`入度（in-degree）` ， 出去的边叫做`出度(out-degree)`。

社交网络中，还会记录两个用户之间的亲密度，联系比较多亲密度就高。每条边都有一个权重(Weight)。这就要用到`带权图（weighted graph）`。

## 图的存储

我们说过，数据的存储最后都是通过两种方式：数组或者链表。那图的存储就有了两种方式：

- 邻接矩阵(Adjacency Matrix)：使用数组
- 邻接表(Adjacency List)： 使用链表

### 邻接矩阵

邻接矩阵使用一个`二维数组`来存储。

- 无向图：若顶点 i 和 j 之间有边，则记 A\[i][j] 和 A\[j][i] 为 1
- 有向图：若顶点 i 指向 顶点 j， 则记 A\[i][j] 为1
- 带权图：数组中存储权重值

<img src="https://static001.geekbang.org/resource/image/62/d2/625e7493b5470e774b5aa91fb4fdb9d2.jpg" alt="img" style="zoom:50%;" />

使用邻接矩阵存储图特别简单，可以快速获取获取两个顶点的关系。但是比较浪费空间，特别是对于社交网络。顶点特别多，但是边很少，就造成了`稀疏图(Sparse Matrix)`。

### 邻接表

邻接表使用链表来存储，如下图所示：

<img src="https://static001.geekbang.org/resource/image/03/94/039bc254b97bd11670cdc4bf2a8e1394.jpg" alt="img" style="zoom:50%;" />

和散列表类似，每个节点对应一个链表，链表中存储的是与这个顶点相连接的顶点。邻接表减少了内存空间，但是操作就相对耗时。例如需要找到节点间的关系，就需要遍历响应顶点的链表。

> 邻接表中，链表可以换成平衡二叉树、散列表或者跳表等来提高操作的效率。