package tree

func deleteNode(root *TreeNode, key int) *TreeNode {

	return nil
}

/**
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
中序查找
前序遍历存数组
重建二叉树？



root 有左右子树，这时可以将 \textit{root}root 的后继节点（比 \textit{root}root 大的最小节点，即它的右子树中的最小节点，记为 \textit{successor}successor）作为新的根节点替代 \textit{root}root，并将 \textit{successor}successor 从 \textit{root}root 的右子树中删除，使得在保持有序性的情况下合并左右子树。
简单证明，\textit{successor}successor 位于 \textit{root}root 的右子树中，因此大于 \textit{root}root 的所有左子节点；\textit{successor}successor 是 \textit{root}root 的右子树中的最小节点，因此小于 \textit{root}root 的右子树中的其他节点。以上两点保持了新子树的有序性。
在代码实现上，我们可以先寻找 \textit{successor}successor，再删除它。\textit{successor}successor 是 \textit{root}root 的右子树中的最小节点，可以先找到 \textit{root}root 的右子节点，再不停地往左子节点寻找，直到找到一个不存在左子节点的节点，这个节点即为 \textit{successor}successor。然后递归地在 \textit{root.right}root.right 调用 \textit{deleteNode}deleteNode 来删除 \textit{successor}successor。因为 \textit{successor}successor 没有左子节点，因此这一步递归调用不会再次步入这一种情况。然后将 \textit{successor}successor 更新为新的 \textit{root}root 并返回。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/delete-node-in-a-bst/solution/shan-chu-er-cha-sou-suo-shu-zhong-de-jie-n6vo/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
