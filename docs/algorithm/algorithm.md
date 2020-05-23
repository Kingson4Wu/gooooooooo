+ Go 刷 Leetcode 系列：恢复二叉搜索树
+ Go 刷 LeetCode 系列：二叉树的最大路径和
+ Go 刷 LeetCode 系列：动态规划（4）分割等和子集（0,1背包问题）


---

### 做题总结
0. 空间换时间 （最小值栈，结构体保存最小值）
1. 递归：（1）自顶向下的递归，（2）自底向上的递归
2. 使用数组代替hashmap的思想（利用数组的下标）, 空间换取时间
2. 使用hashmap存， k:value, v: index (https://leetcode-cn.com/problems/two-sum/solution/liang-shu-zhi-he-by-leetcode-2/)
3. 移位，与或等运算 （n & (n-1) 每次消除一个1.直到n＝0，就能算出多少个1）；与特殊的值与，特别是题目上有提示多少位的整数
4. 边界注意：负数；是否为空；第一个，最后一个；长度是1；不等于0有，总和有可能由正到负再到正再到负，只有算到最后一个叶子才知道；删除的是头或者尾
5. 树遍历：前序遍历，后序，先序（打印）；深度遍历，广度遍历（搜索）
6. 字母的ASCII码是数字且固定
8. 递归改成迭代，使用栈
9. 结果反转
10. 数学计算；善用数学公式推导
11. 搜索树就是有序的意思（中序）
12. 链表遍历：两个指针，快慢指针;两个指针速度相差1，环形情况一定会相遇（在环形区域）（linked-list-cycle）
13. 动态规划：保存已经计算过的结果，供后续使用
14. 树：left和right都为空才是叶子节点
15. 搜索树倒序，迭代法：右子树要全部入栈，直到右子数为空
处理完当前结点，再处理左结点，重复左结点的右子树的入栈操作
（er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof）


---

### 二叉树
+ 二叉树遍历(先序、中序、后序)
  - 先序（中左右），中序（左中右），后序（左右中）
  - 递归遍历和非递归遍历
  - 二叉树遍历(先序、中序、后序):<https://www.jianshu.com/p/456af5480cee e>

### 图

#### 有向图
+ 入度表 indegrees（key为目标顶点, value为源顶点数值）
+ 邻接表 adjacency (key为源顶点, value为目标顶点列表）
+ DFS，递归！
+ 课程表（拓扑排序：入度表BFS法 / DFS法，清晰图解）: https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/

