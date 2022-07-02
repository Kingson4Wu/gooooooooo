package tree

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	start := root

	for start != nil {
		callStart := start
		start = nil

		var pre *Node

		for callStart != nil {

			if callStart.Left != nil {
				if pre == nil {
					pre = callStart.Left
					start = pre
				} else {
					pre.Next = callStart.Left
					pre = callStart.Left
				}
			}

			if callStart.Right != nil {
				if pre == nil {
					pre = callStart.Right
					start = pre
				} else {
					pre.Next = callStart.Right
					pre = callStart.Right
				}
			}

			callStart = callStart.Next
		}

	}

	return root
}

/**

自己做的！

执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
6 MB
, 在所有 Go 提交中击败了
100.00%
的用户
通过测试用例：
55 / 55
*/
