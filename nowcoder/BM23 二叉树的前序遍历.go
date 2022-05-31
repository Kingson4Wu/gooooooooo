package nowcoder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */

var gggggggggggg = []int{}

func preorderTraversal(root *TreeNode) []int {
	// write code here

	call(root)
	return gggggggggggg
}

func call(root *TreeNode) {

	if root == nil {
		return
	}

	gggggggggggg = append(gggggggggggg, root.Val)
	call(root.Left)
	call(root.Right)
}

/**

运行时间：4ms
超过52.36% 用Go提交的代码
占用内存：1176KB
超过21.73%用Go提交的代码

*/

//不用递归怎么写呢？？？！！
