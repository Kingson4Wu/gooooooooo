package tree

/**
112. 路径总和
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，
      		  5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	val := nodeSum(root, sum)

	if val == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, val) || hasPathSum(root.Right, val)
}

func nodeSum(node *TreeNode, sum int) int {
	return sum - node.Val
}

/**
[-2,null,-3]
-5
[5,4,8,11,null,13,4,7,2,null,null,null,1]
22
[1,-2,-3,1,3,-2,null,-1]
-1

1. 没考虑负数！
2. 不等于0有，总和有可能由正到负再到正再到负，只有算到最后一个叶子才知道

递归是自己写的

非递归：（应该是BFS）
class Solution {
  public boolean hasPathSum(TreeNode root, int sum) {
    if (root == null)
      return false;

    LinkedList<TreeNode> node_stack = new LinkedList();
    LinkedList<Integer> sum_stack = new LinkedList();
    node_stack.add(root);
    sum_stack.add(sum - root.val);

    TreeNode node;
    int curr_sum;
    while ( !node_stack.isEmpty() ) {
      node = node_stack.pollLast();
      curr_sum = sum_stack.pollLast();
      if ((node.right == null) && (node.left == null) && (curr_sum == 0))
        return true;

      if (node.right != null) {
        node_stack.add(node.right);
        sum_stack.add(curr_sum - node.right.val);
      }
      if (node.left != null) {
        node_stack.add(node.left);
        sum_stack.add(curr_sum - node.left.val);
      }
    }
    return false;
  }
}

作者：LeetCode
链接：https://leetcode-cn.com/problems/path-sum/solution/lu-jing-zong-he-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
