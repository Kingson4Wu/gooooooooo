package nowcoder

func longestCommonPrefix(strs []string) string {
	// write code here

	total := len(strs)

	if total == 0 {
		return ""
	}

	if total == 1 {
		return strs[0]
	}

	if len(strs[0]) == 0 {
		return ""
	}

	index := 0
	end := 0
	for {

		if len(strs[0]) <= index {
			index--
			end = index
			break
		}
		ch := strs[0][index]
		ok := true

		for i := 0; i < total; i++ {

			if len(strs[i]) <= index {
				ok = false
				index--
				end = index
				break
			}

			if strs[i][index] != ch {
				ok = false
				index--
				end = index
				break
			}
		}
		if !ok {
			break
		} else {
			end = index
		}

		index++
	}

	return strs[0][0 : end+1]
}

/**
描述
给你一个大小为 n 的字符串数组 strs ，其中包含n个字符串 , 编写一个函数来查找字符串数组中的最长公共前缀，返回这个公共前缀。

数据范围： 0 \le n \le 50000≤n≤5000， 0 \le len(strs_i) \le 50000≤len(strs
i
​
 )≤5000
进阶：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
["abca","abc","abca","abc","abcc"]
复制
返回值：
"abc"
复制
示例2
输入：
["abc"]
复制
返回值：
"abc"
*/

/**

内存超限
内存超限:您的程序使用了超过限制的内存

*/

/**
运行时间：63ms
超过26.77% 用Go提交的代码
占用内存：16812KB
超过23.41%用Go提交的代码
*/

/**

代码写得好复杂

答案：
import java.util.*;
public class Solution {
    public String longestCommonPrefix (String[] strs) {
        int n = strs.length;
        //空字符串数组
        if(n == 0)
            return "";
        //遍历第一个字符串的长度
        for(int i = 0; i < strs[0].length(); i++){
            char temp = strs[0].charAt(i);
            //遍历后续的字符串
            for(int j = 1; j < n; j++)
                //比较每个字符串该位置是否和第一个相同
                if(i == strs[j].length() || strs[j].charAt(i) != temp)
                    //不相同则结束
                    return strs[0].substring(0, i);
        }
        //后续字符串有整个字一个字符串的前缀
        return strs[0];
    }
}


*/
