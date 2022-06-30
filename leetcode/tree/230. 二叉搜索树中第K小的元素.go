package tree

var k = 0

func kthSmallest(root *TreeNode, k int) int {

	k = 0
	return kthSmallest2(root, k)
}

func kthSmallest2(root *TreeNode, k int) int {

	if root.Left != nil {
		val := kthSmallest2(root.Left, k)

		if val > 0 {
			return val
		}
	}

	k++
	if k == k {
		return root.Val
	}

	if root.Right != nil {
		val := kthSmallest2(root.Right, k)
		if val > 0 {
			return val
		}
	}

	return 0
}

/**
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
*/

/**

自己做的，中序递归即可，但感觉不是很简洁

官方答案
非递归的中序，使用栈，需要记录状态提前结束遍历的时候，使用栈的方式遍历，代码不会那么丑

func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    for {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        stack, root = stack[:len(stack)-1], stack[len(stack)-1]
        k--
        if k == 0 {
            return root.Val
        }
        root = root.Right
    }
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/kth-smallest-element-in-a-bst/solution/er-cha-sou-suo-shu-zhong-di-kxiao-de-yua-8o07/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：记录子树的结点数

思路和算法

在方法一中，我们之所以需要中序遍历前 k 个元素，是因为我们不知道子树的结点数量，不得不通过遍历子树的方式来获知。

因此，我们可以记录下以每个结点为根结点的子树的结点数，并在查找第 k 小的值时，使用如下方法搜索：

令 node 等于根结点，开始搜索。

对当前结点 node 进行如下操作：

如果 node 的左子树的结点数 left 小于 k−1，则第 k 小的元素一定在 node 的右子树中，令 node 等于其的右子结点，k 等于 k−left−1，并继续搜索；
如果 node 的左子树的结点数 left 等于 k−1，则第 k 小的元素即为 nodenode ，结束搜索并返回 node 即可；
如果 node 的左子树的结点数 left 大于 k−1，则第 k 小的元素一定在 node 的左子树中，令 node 等于其左子结点，并继续搜索。
在实现中，我们既可以将以每个结点为根结点的子树的结点数存储在结点中，也可以将其记录在哈希表中。

type MyBst struct {
    root    *TreeNode
    nodeNum map[*TreeNode]int // 统计以每个结点为根结点的子树的结点数，并存储在哈希表中
}

// 统计以 node 为根结点的子树的结点数
func (t *MyBst) countNodeNum(node *TreeNode) int {
    if node == nil {
        return 0
    }
    t.nodeNum[node] = 1 + t.countNodeNum(node.Left) + t.countNodeNum(node.Right)
    return t.nodeNum[node]
}

// 返回二叉搜索树中第 k 小的元素
func (t *MyBst) kthSmallest(k int) int {
    node := t.root
    for {
        leftNodeNum := t.nodeNum[node.Left]
        if leftNodeNum < k-1 {
            node = node.Right
            k -= leftNodeNum + 1
        } else if leftNodeNum == k-1 {
            return node.Val
        } else {
            node = node.Left
        }
    }
}

func kthSmallest(root *TreeNode, k int) int {
    t := &MyBst{root, map[*TreeNode]int{}}
    t.countNodeNum(root)
    return t.kthSmallest(k)
}



执行用时：
8 ms
, 在所有 Go 提交中击败了
79.39%
的用户
内存消耗：
6.1 MB
, 在所有 Go 提交中击败了
96.56%
的用户
通过测试用例：
93 / 93
*/
