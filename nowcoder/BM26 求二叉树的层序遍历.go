package nowcoder

func levelOrder(root *TreeNode) [][]int {
	// write code here

	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	result = append(result, []int{root.Val})

	for len(queue) > 0 {

		line := []int{}
		newQueue := []*TreeNode{}

		for _, node := range queue {
			if node.Left != nil {
				line = append(line, node.Left.Val)
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				line = append(line, node.Right.Val)
				newQueue = append(newQueue, node.Right)
			}
		}
		if len(line) > 0 {
			result = append(result, line)
		}

		queue = newQueue
	}

	return result
}

/**
给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
例如：
给定的二叉树是{3,9,20,#,#,15,7},
*/

/**

比较简单，使用队列即可
递归总是很难想出来

运行时间：1008ms
超过92.86% 用Go提交的代码
占用内存：22944KB
超过92.86%用Go提交的代码


递归

方法二：递归（扩展思路）
知识点：二叉树递归

递归是一个过程或函数在其定义或说明中有直接或间接调用自身的一种方法，它通常把一个大型复杂的问题层层转化为一个与原问题相似的规模较小的问题来求解。因此递归过程，最重要的就是查看能不能讲原本的问题分解为更小的子问题，这是使用递归的关键。

而二叉树的递归，则是将某个节点的左子树、右子树看成一颗完整的树，那么对于子树的访问或者操作就是对于原树的访问或者操作的子问题，因此可以自我调用函数不断进入子树。

思路：

既然二叉树的前序、中序、后序遍历都可以轻松用递归实现，树型结构本来就是递归喜欢的形式，那我们的层次遍历是不是也可以尝试用递归来试试呢？

按行遍历的关键是每一行的深度对应了它输出在二维数组中的深度，即深度可以与二维数组的下标对应，那我们可以在递归的访问每个节点的时候记录深度：

void traverse(TreeNode root, int depth)
进入子节点则深度加1：

//递归左右时深度记得加1
traverse(root.left, depth + 1);
traverse(root.right, depth + 1);
每个节点值放入二维数组相应行。

res[depth - 1].push_back(root->val);
因此可以用递归实现：

终止条件： 遍历到了空节点，就不再继续，返回。
返回值： 将加入的输出数组中的结果往上返回。
本级任务： 处理按照上述思路处理非空节点，并进入该节点的子节点作为子问题。
具体做法：

step 1：首先判断二叉树是否为空，空树没有遍历结果。
step 2：使用递归进行层次遍历输出，每次递归记录当前二叉树的深度，每当遍历到一个节点，如果为空直接返回。
step 3：如果遍历的节点不为空，输出二维数组中一维数组的个数（即代表了输出的行数）小于深度，说明这个节点应该是新的一层，我们在二维数组中增加一个一维数组，然后再加入二叉树元素。
step 4：如果不是step 3的情况说明这个深度我们已经有了数组，直接根据深度作为下标取出数组，将元素加在最后就可以了。
step 5：处理完这个节点，再依次递归进入左右节点，同时深度增加。因为我们进入递归的时候是先左后右，那么遍历的时候也是先左后右，正好是层次遍历的顺序。


import java.util.*;
public class Solution {
    //记录输出
    ArrayList<ArrayList<Integer> > res = new ArrayList();
    void traverse(TreeNode root, int depth) {
        if(root != null){
            //新的一层
            if(res.size() < depth){
                ArrayList<Integer> row = new ArrayList();
                res.add(row);
                row.add(root.val);
            //读取该层的一维数组，将元素加入末尾
            }else{
                ArrayList<Integer> row = res.get(depth - 1);
                row.add(root.val);
            }
        }
        else
            return;
        //递归左右时深度记得加1
        traverse(root.left, depth + 1);
        traverse(root.right, depth + 1);
    }
    public ArrayList<ArrayList<Integer>> levelOrder (TreeNode root) {
        if(root == null)
            //如果是空，则直接返回
            return res;
        //递归层次遍历
        traverse(root, 1);
        return res;
    }
}

*/
