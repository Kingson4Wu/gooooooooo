package tree

import "strconv"

func PrintTree(root *TreeNode) [][]string {

	if root == nil {
		return [][]string{}
	}
	result := [][]*TreeNode{}

	hasNext := true
	lineNum := 0
	line := []*TreeNode{root}
	result = append(result, line)
	for hasNext {
		hasNext = false
		nextLineCount := 1<<(lineNum+2) - 1
		lineResult := make([]*TreeNode, nextLineCount)
		j := 0
		for i := 0; i < len(result[lineNum]); i++ {
			node := result[lineNum][i]
			if node == nil {
				lineResult[j] = nil
				j++
				lineResult[j] = nil
				j++
				lineResult[j] = nil
				j += 2
			} else {
				lineResult[j] = node.Left
				if node.Left != nil {
					hasNext = true
				}
				j++
				lineResult[j] = nil
				j++
				lineResult[j] = node.Right
				if node.Right != nil {
					hasNext = true
				}
				j += 2
			}
			i++
			//循环那里已经会在i++
		}
		if hasNext {

			newResult := [][]*TreeNode{}
			for _, nodeArr := range result {
				lineResult := []*TreeNode{}
				for _, node := range nodeArr {
					lineResult = append(lineResult, nil)
					lineResult = append(lineResult, node)
				}
				lineResult = append(lineResult, nil)

				newResult = append(newResult, lineResult)
			}
			newResult = append(newResult, lineResult)
			result = newResult
		}
		lineNum++

	}

	valueResult := [][]string{}
	for _, nodeArr := range result {
		lineResult := []string{}
		for _, node := range nodeArr {
			if node != nil {
				lineResult = append(lineResult, strconv.Itoa(node.Val))
			} else {
				lineResult = append(lineResult, "")
			}

		}
		valueResult = append(valueResult, lineResult)
	}

	return valueResult

}

/**
貌似会写，但是没写对，结果还跟狗屎一样！

输入：
[5,3,6,2,4,null,7]
输出：
[["","","","5","","",""],["","3","","","","6",""],["2","","4","","","",""]]
预期结果：
[["","","","5","","",""],["","3","","","","6",""],["2","","4","","","","7"]]


自己写的恶心代码，空间换时间，一次层次遍历

算好各种下标
规律
增加下一行，前面的每一行的元素加一个间隔


执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
2.5 MB
, 在所有 Go 提交中击败了
6.82%
的用户
通过测试用例：
73 / 73

*/
