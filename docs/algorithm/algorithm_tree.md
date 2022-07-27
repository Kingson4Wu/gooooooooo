+ 树遍历：前序遍历，后序，先序（打印）；深度遍历，广度遍历（搜索）
+ 搜索树就是有序的意思（中序）- 二叉搜索树满足每个节点的左子树上的所有节点均严格小于当前节点且右子树上的所有节点均严格大于当前节点。（注意不能等于！！！！）
+ 树：left和right都为空才是叶子节点
+ 搜索树倒序，迭代法：右子树要全部入栈，直到右子数为空
处理完当前结点，再处理左结点，重复左结点的右子树的入栈操作
（er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof）
+ 完全二叉树--若二叉树的深度为 h，除第 h 层外，其它各层的结点数都达到最大个数，第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。（第 h 层可能包含 [1~2h] 个节点）
判断是否完全二叉树：广度优先搜索，队列，判断是不是连续两个不存在
nowcoder/BM35 判断是不是完全二叉树.go
+ 平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树
+ 堆的结构可以分为大顶堆和小顶堆，是一个完全二叉树


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
        //找到 - 在子结点找到了
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

+ 根据前序和中序重建二叉树
（1）前序第一个为根节点
（2）找到中序中根节点的位置，分成两组，分治递归构造二叉树

前序和中序的特点：
前序=【左子树】+【根节点】+【右子树】
中序=【根节点】+【左子树】+【右子树】

+ nowcoder/BM39 序列化二叉树.go
(1)一个字符串的前序遍历（标识节点分割和空节点），可以重建二叉树（递归方式）
(2)一个字符串的前序遍历和一个字符串的中序遍历（标识节点分割），可以重建二叉树（递归方式）

+ nowcoder/BM41 输出二叉树的右视图.go -- 层次遍历！！！
请根据二叉树的前序遍历，中序遍历恢复二叉树，并打印出二叉树的右视图
构建完成二叉树之后，我们只需要对二叉树进行层次遍历。遍历到每一层时，将该层的最右边节点加入数组即可。
 
 int size = queue.size();
 for(int i = 0; i < size; i++){
    if(i == size - 1) result.add(temp.val);
 }

```java
 public ArrayList<Integer> right(TreeNode root){
        ArrayList<Integer> result = new ArrayList<Integer>();
        if(root == null)
            return result;
        // 层次遍历
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        queue.offer(root);
        while(!queue.isEmpty()){
            int size = queue.size();
            for(int i = 0; i < size; i++){
                TreeNode temp = queue.poll();
                // 记录每一层的最右边节点
                if(i == size - 1) result.add(temp.val);
                if(temp.left != null) queue.offer(temp.left);
                if(temp.right != null) queue.offer(temp.right);
            }
        }
        return result;
    }
/**
————————————————
版权声明：本文为CSDN博主「木水先生」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_39963192/article/details/115119704
*/

```

+ 非递归的中序，使用栈，需要记录状态提前结束遍历的时候，使用栈的方式遍历，代码不会那么丑
leetcode/tree/230. 二叉搜索树中第K小的元素.go

+ leetcode/tree/95. 不同的二叉搜索树 II.go
回溯递归！！！

```go
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return nil
    }
    return helper(1, n)
}

func helper(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    allTrees := []*TreeNode{}
    // 枚举可行根节点
    for i := start; i <= end; i++ {
        // 获得所有可行的左子树集合
        leftTrees := helper(start, i - 1)
        // 获得所有可行的右子树集合
        rightTrees := helper(i + 1, end)
        // 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
        for _, left := range leftTrees {
            for _, right := range rightTrees {
                currTree := &TreeNode{i, nil, nil}
                currTree.Left = left
                currTree.Right = right
                allTrees = append(allTrees, currTree)
            }
        }
    }
    return allTrees
}

```

---
# 搜索
+ 树的遍历方法有广度优先（层序遍历），以及深度优先两种方法，分成先序遍历，中序遍历，后序遍历三种。
+ 深度优先使用栈，广度优先使用队列
## 深度优先搜索
+ 递归地实现 DFS 时，似乎不需要使用任何栈。但实际上，我们使用的是由系统提供的隐式栈，也称为调用栈（Call Stack）
### 栈实现

+ 先序
(1)`root!=nil || len(stack)>0`
(2) root!=nil 处理逻辑，入栈，左子树赋值
(3) else 出栈，右子树赋值
```go
func (root *TreeNode) preorder() []int{       //非递归前序遍历
	res:=[]int{}
	if root==nil{
		return res
	}
	stack:=[]*TreeNode{}           //定义一个栈储存节点信息
	for root!=nil || len(stack)!=0{
		if root!=nil{
			res=append(res,root.data)
			stack=append(stack,root)        
			root=root.Lchild
		}else{
			root=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			root=root.Rchild
		}
	}
	return res
}

```

+ 中序

(1)`root!=nil || len(stack)>0`
(2) root!=nil 入栈，左子树赋值
(3) else 出栈，处理逻辑，右子树赋值


(1)`root!=nil || len(stack)>0`
(2)左子树循环入栈
(3)出栈，处理逻辑
(4)`root=root.Right`

``` go
func (root *TreeNode) inorder()[]int{
	res:=[]int{}
	if root==nil{
		return res
	}
	stack:=[]*TreeNode{}
	for root!=nil || len(stack)!=0{
		if root!=nil{
			stack=append(stack,root)
			root=root.Lchild
		}else{
			root=stack[len(stack)-1]
			res=append(res,root.data)
			stack=stack[:len(stack)-1]
			root=root.Rchild
		}
	}
	return res
}
```

```go
func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    inorder := math.MinInt64
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    return true
}

```

+ 后序
有点复杂，日

```go
func (root *TreeNode)postorder() []int {
	res:=[]int{}
	if root==nil{
		return res
	}
	stack:=[]*TreeNode{}
	pre:=&TreeNode{}
	stack=append(stack,root)
	for len(stack)!=0{
		cur:=stack[len(stack)-1]
		if (cur.Lchild==nil && cur.Rchild==nil) || (pre!=nil &&(pre==cur.Lchild || pre==cur.Rchild)){
			res=append(res,cur.data)
			pre=cur
			stack=stack[:len(stack)-1]
		}else{
			if cur.Rchild!=nil{
				stack=append(stack,cur.Rchild)
			}
			if cur.Lchild!=nil{
				stack=append(stack,cur.Lchild)
			}
		}
	}
	return res
}

```


递归
```go

func findBottomLeftValue(root *TreeNode) (curVal int) {
    curHeight := 0
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, height int) {
        if node == nil {
            return
        }
        height++
        dfs(node.Left, height)
        dfs(node.Right, height)
        if height > curHeight {
            curHeight = height
            curVal = node.Val
        }
    }
    dfs(root, 0)
    return
}

```

+ 广度优先（层序遍历）
```java
void LevelOrderTraversal( BinTree BT)
{
    Queue Q;
    BinTree T;
    if(!BT) return;  //若是空树则直接返回
    Q = CreatQueue(maxsize); //创建并初始化队列
    AddQ(Q,BT);
    while(!IsEmpty(s))
    {
        T=deleteQ(Q);     //出队
        printf("%d\n",T->data);  //访问该节点
        if(T->Left) AddQ(Q,T->Left);   //分别将其左右子入队
        if(T->Right) AddQ(Q,T->Right);  
    }
}

```

+ leetcode/tree/669. 修剪二叉搜索树.go
```java
class Solution {
    public TreeNode trimBST(TreeNode root, int L, int R) {
        if (root == null) return root;
        if (root.val > R) return trimBST(root.left, L, R);
        if (root.val < L) return trimBST(root.right, L, R);

        root.left = trimBST(root.left, L, R);
        root.right = trimBST(root.right, L, R);
        return root;
    }
}

```

+ leetcode/tree/114. 二叉树展开为链表.go
前两种方法都借助前序遍历，前序遍历过程中需要使用栈存储节点。有没有空间复杂度是 O(1)O(1) 的做法呢？

注意到前序遍历访问各节点的顺序是根节点、左子树、右子树。如果一个节点的左子节点为空，则该节点不需要进行展开操作。
如果一个节点的左子节点不为空，则该节点的左子树中的最后一个节点被访问之后，该节点的右子节点被访问。
该节点的左子树中最后一个被访问的节点是左子树中的最右边的节点，也是该节点的前驱节点。
因此，问题转化成寻找当前节点的前驱节点。

具体做法是，对于当前节点，如果其左子节点不为空，则在其左子树中找到最右边的节点，作为前驱节点，
将当前节点的右子节点赋给前驱节点的右子节点，然后将当前节点的左子节点赋给当前节点的右子节点，
并将当前节点的左子节点设为空。对当前节点处理结束后，继续处理链表中的下一个节点，直到所有节点都处理结束。
```go

func flatten(root *TreeNode)  {
    curr := root
    for curr != nil {
        if curr.Left != nil {
            next := curr.Left
            predecessor := next
            for predecessor.Right != nil {
                predecessor = predecessor.Right
            }
            predecessor.Right = curr.Right
            curr.Left, curr.Right = nil, next
        }
        curr = curr.Right
    }
}

```

+ leetcode/tree/117. 填充每个节点的下一个右侧节点指针 II.go
方法二：使用已建立的 next 指针 ！！！！
空间复杂度：O(1)，不需要存储额外的节点。

+ leetcode/tree/450. 删除二叉搜索树中的节点.go

root 为叶子节点，没有子树。此时可以直接将它删除，即返回空。
root 只有左子树，没有右子树。此时可以将它的左子树作为新的子树，返回它的左子节点。
root 只有右子树，没有左子树。此时可以将它的右子树作为新的子树，返回它的右子节点。

root 有左右子树，这时可以将 root 的后继节点（比 root 大的最小节点，即它的右子树中的最小节点，记为 successor作为新的根节点替代 root，并将 successor 从 root 的右子树中删除，使得在保持有序性的情况下合并左右子树。
简单证明，successor 位于 root 的右子树中，因此大于 root 的所有左子节点；successor 是 root 的右子树中的最小节点，因此小于 root 的右子树中的其他节点。以上两点保持了新子树的有序性。

+ leetcode/tree/508. 出现次数最多的子树元素和.go

```go
func findFrequentTreeSum(root *TreeNode) (ans []int) {
    cnt := map[int]int{}
    maxCnt := 0
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        sum := node.Val + dfs(node.Left) + dfs(node.Right)
        cnt[sum]++
        if cnt[sum] > maxCnt {
            maxCnt = cnt[sum]
        }
        return sum
    }
    dfs(root)

    for s, c := range cnt {
        if c == maxCnt {
            ans = append(ans, s)
        }
    }
    return
}

```
+ 匿名函数写法！！


+ leetcode/tree/863. 二叉树中所有距离为 K 的结点.go
给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。
返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。

深度优先搜索 + 哈希表
若将 target 当作树的根结点，我们就能从 target 出发，使用深度优先搜索去寻找与 target 距离为 kk 的所有结点，即深度为 kk 的所有结点。

由于输入的二叉树没有记录父结点，为此，我们从根结点 root 出发，使用深度优先搜索遍历整棵树，同时用一个哈希表记录每个结点的父结点。

然后从 target 出发，使用深度优先搜索遍历整棵树，除了搜索左右儿子外，还可以顺着父结点向上搜索。

代码实现时，由于每个结点值都是唯一的，哈希表的键可以用结点值代替。此外，为避免在深度优先搜索时重复访问结点，递归时额外传入来源结点 from，在递归前比较目标结点是否与来源结点相同，不同的情况下才进行递归。
(左右子树是往下的，父节点是往上的，所以可能相互重复)


```go

func distanceK(root, target *TreeNode, k int) (ans []int) {
    // 从 root 出发 DFS，记录每个结点的父结点
    parents := map[int]*TreeNode{}
    var findParents func(*TreeNode)
    findParents = func(node *TreeNode) {
        if node.Left != nil {
            parents[node.Left.Val] = node
            findParents(node.Left)
        }
        if node.Right != nil {
            parents[node.Right.Val] = node
            findParents(node.Right)
        }
    }
    findParents(root)

    // 从 target 出发 DFS，寻找所有深度为 k 的结点
    var findAns func(*TreeNode, *TreeNode, int)
    findAns = func(node, from *TreeNode, depth int) {
        if node == nil {
            return
        }
        if depth == k { // 将所有深度为 k 的结点的值计入结果
            ans = append(ans, node.Val)
            return
        }
        if node.Left != from {
            findAns(node.Left, node, depth+1)
        }
        if node.Right != from {
            findAns(node.Right, node, depth+1)
        }
        if parents[node.Val] != from {
            findAns(parents[node.Val], node, depth+1)
        }
    }
    findAns(target, nil, 0)
    return
}

```