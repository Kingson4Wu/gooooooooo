+ 树遍历：前序遍历，后序，先序（打印）；深度遍历，广度遍历（搜索）
+ 搜索树就是有序的意思（中序）- 二叉搜索树满足每个节点的左子树上的所有节点均严格小于当前节点且右子树上的所有节点均严格大于当前节点。
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
``` java
void InOrderTraversal( BinTree BT)
{
    BinTree T=BT;
    stack s=CreatStack(maxsize);     //定义一个栈
    while(T||!IsEmpty(s))           
    {
        while(T){
            push(s,T);                  //一直向左将节点入栈，直到左子树为空。
            T=T->Left;
        }
    if(!IsEmpty(s)){                            
        T=pop(s);                        //节点出栈                          
        printf("%5d",T->Data);           //输出节点
        T=T->Right;                      //再转向右子树。
    }
    }
}

```
+ 中序
``` java
void InOrderTraversal( BinTree BT)
{
    BinTree T=BT;
    stack s=CreatStack(maxsize);
    while(T||!IsEmpty(s))
    {
        while(T){
            push(s,T);
            printf("%5d",T->Data);
            t=t->Left;
        }
    if(!IsEmpty(s)){
        T=pop(s);
        T=T->Right;
    }
    }
}

```
+ 后序
```java
void PostOrderTraversal(Bintree BT) {  //给节点增加访问次数的属性Visit，初始化为0
    Bintree T BT;
    Stack S = CreateStack(Maxsize);
    while (T || !IsEmpty(S)) {
        while (T) {
            if (T->Visit == 0) {//虽然没必要判断，为便于理解
                T->Visit++;
                Push(S, T);  //第一次入栈，不访问
            }
            T = T->left;   //转向左子树
        }
        if (!IsEmpty(S)) {
            T = Pop(s);
            if (T->Visit == 2)    {
                printf("%d", T->Data);//第三次碰到它，访问节点，可以彻底从堆栈弹出了
                T = NULL;//左右子数均已经访问过
            }
            else {
                T->Visit++;
                Push(S, T);  //第二次入栈，不访问，（相当于T没有出栈）
                T = T->Right;  //转向右子树
            }
        }
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