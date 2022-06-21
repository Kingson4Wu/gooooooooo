package nowcoder

var indexMap = make(map[int]int)
var sortValue = 0

func lowestCommonAncestor2(root *TreeNode, o1 int, o2 int) int {
	// write code here

	buildSearchTree(root)

	o1Value := indexMap[o1]
	o2Value := indexMap[o2]

	return lowestCommonAncestor3(root, o1Value, o2Value)
}

func lowestCommonAncestor3(root *TreeNode, o1 int, o2 int) int {

	if root == nil {
		return -1
	}

	rootValue := indexMap[root.Val]

	if o1 >= rootValue && o2 <= rootValue || o1 <= rootValue && o2 >= rootValue {
		return root.Val
	} else if o1 <= rootValue && o2 <= rootValue {
		return lowestCommonAncestor3(root.Left, o1, o2)
	} else {
		return lowestCommonAncestor3(root.Right, o1, o2)
	}
}

func buildSearchTree(root *TreeNode) {

	if root == nil {
		return
	}

	buildSearchTree(root.Left)

	indexMap[root.Val] = sortValue
	sortValue++

	buildSearchTree(root.Right)
}

/**
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。

本题保证二叉树中每个节点的val值均不相同。
*/

/**

个人想法，前序遍历，构造二叉搜索树，将val和下标存到map。
通过nowcoder/BM37 二叉搜索树的最近公共祖先.go的方法判断最近公共祖先

运行时间：133ms
超过0.00% 用Go提交的代码
占用内存：29348KB
超过0.00%用Go提交的代码

*/

/**

方法一：路径比较法(推荐使用)
知识点：深度优先搜索（dfs）

深度优先搜索一般用于树或者图的遍历，其他有分支的（如二维矩阵）也适用。它的原理是从初始点开始，一直沿着同一个分支遍历，直到该分支结束，然后回溯到上一级继续沿着一个分支走到底，如此往复，直到所有的节点都有被访问到。

思路：

既然要找到二叉树中两个节点的最近公共祖先，那我们可以考虑先找到两个节点全部祖先，可以得到从根节点到目标节点的路径，然后依次比较路径得出谁是最近的祖先。

找到两个节点的所在可以深度优先搜索遍历二叉树所有节点进行查找。

具体做法：

step 1：利用dfs求得根节点到两个目标节点的路径：每次选择二叉树的一棵子树往下找，同时路径数组增加这个遍历的节点值。
step 2：一旦遍历到了叶子节点也没有，则回溯到父节点，寻找其他路径，回溯时要去掉数组中刚刚加入的元素。
step 3：然后遍历两条路径数组，依次比较元素值。
step 4：找到两条路径第一个不相同的节点即是最近公共祖先。

import java.util.*;
public class Solution {
    //记录是否找到到o的路径
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
    public int lowestCommonAncestor (TreeNode root, int o1, int o2) {
        ArrayList<Integer> path1 = new ArrayList<Integer>();
        ArrayList<Integer> path2 = new ArrayList<Integer>();
        //求根节点到两个节点的路径
        dfs(root, path1, o1);
        //重置flag，查找下一个
        flag = false;
        dfs(root, path2, o2);
        int res = 0;
        //比较两个路径，找到第一个不同的点
        for(int i = 0; i < path1.size() && i < path2.size(); i++){
            int x = path1.get(i);
            int y = path2.get(i);
            if(x == y)
                //最后一个相同的节点就是最近公共祖先
                res = x;
            else
                break;
        }
        return res;
    }
}

方法二：递归（扩展思路）
知识点：二叉树递归

思路：

我们也可以讨论几种情况：

step 1：如果o1和o2中的任一个和root匹配，那么root就是最近公共祖先。
step 2：如果都不匹配，则分别递归左、右子树。
step 3：如果有一个节点出现在左子树，并且另一个节点出现在右子树，则root就是最近公共祖先.
step 4：如果两个节点都出现在左子树，则说明最低公共祖先在左子树中，否则在右子树。
step 5：继续递归左、右子树，直到遇到step1或者step3的情况。

import java.util.*;
public class Solution {
    public int lowestCommonAncestor (TreeNode root, int o1, int o2) {
        //该子树没找到，返回-1
        if(root == null)
            return -1;
        //该节点是其中某一个节点
        if(root.val == o1 || root.val == o2)
            return root.val;
        //左子树寻找公共祖先
        int left = lowestCommonAncestor(root.left, o1, o2);
        //右子树寻找公共祖先
        int right = lowestCommonAncestor(root.right, o1, o2);
        //左子树为没找到，则在右子树中
        if(left == -1)
            return right;
        //右子树没找到，则在左子树中
        if(right == -1)
            return left;
        //否则是当前节点
        return root.val;
    }
}



*/
