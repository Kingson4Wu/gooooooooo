package tree

func isValidSerialization(preorder string) bool {

	for _, c := range preorder {

		if c == '#' {

		}
	}

	return false
}

/**
序列化二叉树的一种方法是使用 前序遍历 。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
*/

/**

方法一：栈
我们可以定义一个概念，叫做槽位。一个槽位可以被看作「当前二叉树中正在等待被节点填充」的那些位置。

二叉树的建立也伴随着槽位数量的变化。每当遇到一个节点时：

如果遇到了空节点，则要消耗一个槽位；

如果遇到了非空节点，则除了消耗一个槽位外，还要再补充两个槽位。

此外，还需要将根节点作为特殊情况处理。

我们使用栈来维护槽位的变化。栈中的每个元素，代表了对应节点处剩余槽位的数量，而栈顶元素就对应着下一步可用的槽位数量。当遇到空节点时，仅将栈顶元素减 11；当遇到非空节点时，将栈顶元素减 11 后，再向栈中压入一个 22。无论何时，如果栈顶元素变为 00，就立刻将栈顶弹出。

遍历结束后，若栈为空，说明没有待填充的槽位，因此是一个合法序列；否则若栈不为空，则序列不合法。此外，在遍历的过程中，若槽位数量不足，则序列不合法。

func isValidSerialization(preorder string) bool {
    n := len(preorder)
    stk := []int{1}
    for i := 0; i < n; {
        if len(stk) == 0 {
            return false
        }
        if preorder[i] == ',' {
            i++
        } else if preorder[i] == '#' {
            stk[len(stk)-1]--
            if stk[len(stk)-1] == 0 {
                stk = stk[:len(stk)-1]
            }
            i++
        } else {
            // 读一个数字
            for i < n && preorder[i] != ',' {
                i++
            }
            stk[len(stk)-1]--
            if stk[len(stk)-1] == 0 {
                stk = stk[:len(stk)-1]
            }
            stk = append(stk, 2)
        }
    }
    return len(stk) == 0
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/solution/yan-zheng-er-cha-shu-de-qian-xu-xu-lie-h-jghn/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：计数
能否将方法一的空间复杂度优化至 O(1)O(1) 呢？

回顾方法一的逻辑，如果把栈中元素看成一个整体，即所有剩余槽位的数量，也能维护槽位的变化。

因此，我们可以只维护一个计数器，代表栈中所有元素之和，其余的操作逻辑均可以保持不变。

func isValidSerialization(preorder string) bool {
    n := len(preorder)
    slots := 1
    for i := 0; i < n; {
        if slots == 0 {
            return false
        }
        if preorder[i] == ',' {
            i++
        } else if preorder[i] == '#' {
            slots--
            i++
        } else {
            // 读一个数字
            for i < n && preorder[i] != ',' {
                i++
            }
            slots++ // slots = slots - 1 + 2
        }
    }
    return slots == 0
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/solution/yan-zheng-er-cha-shu-de-qian-xu-xu-lie-h-jghn/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
