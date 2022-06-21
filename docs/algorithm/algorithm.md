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
8. 递归改成迭代，使用栈 !!!!!
9. 结果反转；涉及反转的可以考虑使用辅助栈
10. 数学计算；善用数学公式推导
11. 搜索树就是有序的意思（中序）- 二叉搜索树满足每个节点的左子树上的所有节点均严格小于当前节点且右子树上的所有节点均严格大于当前节点。
12. 链表遍历：两个指针，快慢指针;两个指针速度相差1，环形情况一定会相遇（在环形区域）（linked-list-cycle）
因为快慢指针总会到一个圈里不停循环，极端情况下，快指针在圈内循环多几次，总能追上慢指针
13. 动态规划：保存已经计算过的结果，供后续使用
14. 树：left和right都为空才是叶子节点
15. 搜索树倒序，迭代法：右子树要全部入栈，直到右子数为空
处理完当前结点，再处理左结点，重复左结点的右子树的入栈操作
（er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof）
16. 最大值队列（辅助栈，双端队列）
17. 遇到出现次数的题，考虑原地hash的骚操作，利用转负数保留原来的值同时记录状态。出现双数（两次）的，考虑交替正负赋值
nowcoder/NC30 缺失的第一个正整数.go
18. 完全二叉树--若二叉树的深度为 h，除第 h 层外，其它各层的结点数都达到最大个数，第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。（第 h 层可能包含 [1~2h] 个节点）
判断是否完全二叉树：广度优先搜索，队列，判断是不是连续两个不存在
nowcoder/BM35 判断是不是完全二叉树.go
19. 平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。
20. 旋转数组
三次翻转 ！！！！
21. 最长无重复子数组
滑动窗口（双指针）+ hash表！！！！

---

### 二叉树
+ 二叉树遍历(先序、中序、后序)
  - 先序（中左右），中序（左中右），后序（左右中）
  - 递归遍历和非递归遍历
  - 二叉树遍历(先序、中序、后序):<https://www.jianshu.com/p/456af5480cee e>

+ 写代码前，先画图整理好思路！！！！
+ 前序遍历使用递归即可
+ 中序遍历使用栈（有递归解法....）
`nowcoder/BM24 二叉树的中序遍历.go`
递归真的好难理解！！！

```go
   public void inorder(List<Integer> list, TreeNode root){
        //遇到空节点则返回
        if(root == null)
            return;
        //先去左子树
        inorder(list, root.left);
        //再访问根节点
        list.add(root.val);
        //最后去右子树
        inorder(list, root.right);
    }

```
+ 后序遍历同样可以使用递归（和中序递归一样处理）

+ nowcoder/BM31 对称的二叉树.go
还是用递归！！！
```java
public class Solution {
    boolean recursion(TreeNode root1, TreeNode root2){
        //可以两个都为空
        if(root1 == null && root2 == null)
            return true;
        //只有一个为空或者节点值不同，必定不对称
        if(root1 == null || root2 == null || root1.val != root2.val)
            return false;
        //每层对应的节点进入递归比较
        return recursion(root1.left, root2.right) && recursion(root1.right, root2.left);
    }
    boolean isSymmetrical(TreeNode pRoot) {
        return recursion(pRoot, pRoot);
    }
}
```
+ 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先！！！

方法一：搜索路径比较（推荐使用）
step 1：根据二叉搜索树的性质，从根节点开始查找目标节点，当前节点比目标小则进入右子树，当前节点比目标大则进入左子树，直到找到目标节点。这个过程成用数组记录遇到的元素。
step 2：分别在搜索二叉树中找到p和q两个点，并记录各自的路径为数组。
step 3：同时遍历两个数组，比较元素值，最后一个相等的元素就是最近的公共祖先。


方法二：一次遍历（扩展思路）
step 1：首先检查空节点，空树没有公共祖先。
step 2：对于某个节点，比较与p、q的大小，若p、q在该节点两边说明这就是最近公共祖先。
step 3：如果p、q都在该节点的左边，则递归进入左子树。
step 4：如果p、q都在该节点的右边，则递归进入右子树。

//pq在该节点两边说明这就是最近公共祖先(因为从根节点遍历的，那么最先遇到符合条件的就是最近公共祖先！！！！)
        if((p >= root.val && q <= root.val) || (p <= root.val && q >= root.val))
            return root.val;


+ nowcoder/BM38 在二叉树中找到两个节点的最近公共祖先.go
(1)深度优先搜索，回溯法递归找路径，比较两个路径第一个不相同的前一个就是公共祖先
```java
 public boolean flag = false;
    //求得根节点到目标节点的路径
    public void dfs(TreeNode root, ArrayList<Integer> path, int o){
        if(flag || root == null)
            return;
        path.add(root.val);
        //节点值都不同，可以直接用值比较
        if(root.val == o){
            flag = true;
            return;
        }
        //dfs遍历查找
        dfs(root.left, path, o);
        dfs(root.right, path, o);
        //找到
        if(flag)
            return;
        //回溯
        path.remove(path.size() - 1);
    }
```
(2)二叉树递归

思路：

我们也可以讨论几种情况：

step 1：如果o1和o2中的任一个和root匹配，那么root就是最近公共祖先。
step 2：如果都不匹配，则分别递归左、右子树。
step 3：如果有一个节点出现在左子树，并且另一个节点出现在右子树，则root就是最近公共祖先.
step 4：如果两个节点都出现在左子树，则说明最低公共祖先在左子树中，否则在右子树。
step 5：继续递归左、右子树，直到遇到step1或者step3的情况。


### 图

#### 有向图
+ 入度表 indegrees（key为目标顶点, value为源顶点数值）
+ 邻接表 adjacency (key为源顶点, value为目标顶点列表）
+ DFS，递归！
+ 课程表（拓扑排序：入度表BFS法 / DFS法，清晰图解）: https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/

##### DFS
+ TODO
##### BFS
+ TODO

### 堆
+ 堆的结构可以分为大顶堆和小顶堆，是一个完全二叉树
+ 底层数组结构
+ 大顶堆：`arr(i)>arr(2*i+1) && arr(i)>arr(2*i+2)`；小顶堆：`arr(i)<arr(2*i+1) && arr(i)<arr(2*i+2)`
+ 根找左右：左：`(i+1)*2 - 1`，右：`(i+1)*2` ;
+ 左找根右：根：`(i+1)/2 - 1`，右：`i+1`
+ 右照根左：根：`i/2 - 1`, 左：`i-1`

+ 1 初始化堆，比如小堆：加入堆最后一个结点后，与根节点对比，比根节点小则交换，并继续，否则结束完成（自底向上）
+ 2 调整堆（拿出堆中第一个元素后，即根最后一个结点交换后），将根节点和左右两个元素比较，与较小的那个交换，并继续，若都比自己大，或者已经是叶子结点，则结束完成（自上往下）

### 快速排序
+ 高低下标的值交替被替换，最后得到中间下标，把参照值设置回去
+ topk，结果无序排序的时候可以用

### 动态规划
+ 总结公式：
+ 兑换零钱（一）
+ 打家劫舍（一）
此转移方程为dp[i]=max(dp[i−1],nums[i−1]+dp[i−2])
+ 打家劫舍（二）
这一问在第一问的基础上添加了房屋首尾相连的条件，所以首尾两个数字之间我们最多只能选取一个。

∙ \bullet∙ 既然这样我们就可以将整个数组进行分割，分为 [0,n-2] 与 [1,n-1] 两个部分，分别求解它们的最大值然后再选择两个之间的较大值作为最终结果，其余情况与第一问相同。
————————————————
版权声明：本文为CSDN博主「桃陉」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_46308081/article/details/119087704

### 链表
+ 链表的问题基本都挺容易想，但是写起来比价麻烦，需要许多中间变量，所以做之前要先画图整理清楚，否则写的时候很容易乱



----

+ 蓄水池抽样算法（Reservoir Sampling）:<https://www.jianshu.com/p/7a9ea6ece2af>