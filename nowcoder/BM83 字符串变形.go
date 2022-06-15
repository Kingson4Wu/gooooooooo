package nowcoder

func trans(s string, n int) string {
	// write code here

	a := 'a'
	a = []rune("sss")[0]
	if a > 'Z' {
		return ""
	}
	return ""
}

/**
描述
对于一个长度为 n 字符串，我们需要对它做一些变形。

首先这个字符串中包含着一些空格，就像"Hello World"一样，然后我们要做的是把这个字符串中由空格隔开的单词反序，同时反转每个字符的大小写。

比如"Hello World"变形后就变成了"wORLD hELLO"。

数据范围: 1\le n \le 10^61≤n≤10
6
  , 字符串中包括大写英文字母、小写英文字母、空格。
进阶：空间复杂度 O(n) ， 时间复杂度 O(n)
输入描述：
给定一个字符串s以及它的长度n(1 ≤ n ≤ 10^6)
返回值描述：
请返回变形后的字符串。题目保证给定的字符串均由大小写字母和空格构成。
示例1
输入：
"This is a sample",16
复制
返回值：
"SAMPLE A IS tHIS"
复制
示例2
输入：
"nowcoder",8
复制
返回值：
"NOWCODER"
复制
示例3
输入：
"iOS",3
复制
返回值：
"Ios"


方法一：双逆转（推荐使用）
思路：

将单词位置的反转，那肯定前后都是逆序，不如我们先将整个字符串反转，这样是不是单词的位置也就随之反转了。但是单词里面的成分也反转了啊，既然如此我们再将单词里面的部分反转过来就行。

具体做法：

step 1：遍历字符串，遇到小写字母，转换成大写，遇到大写字母，转换成小写，遇到空格正常不变。
step 2：第一次反转整个字符串，这样基本的单词逆序就有了，但是每个单词的字符也是逆的。
step 3：再次遍历字符串，以每个空间为界，将每个单词反转回正常。

import java.util.*;
public class Solution {
    public String trans(String s, int n) {
        if(n==0)
            return s;
        StringBuffer res=new StringBuffer();
        for(int i = 0; i < n; i++){
            //大小写转换
            if(s.charAt(i) <= 'Z' && s.charAt(i) >= 'A')
                res.append((char)(s.charAt(i) - 'A' + 'a'));
            else if(s.charAt(i) >= 'a' && s.charAt(i) <= 'z')
                res.append((char)(s.charAt(i) - 'a' + 'A'));
            else
                //空格直接复制
                res.append(s.charAt(i));
        }
        //翻转整个字符串
        res = res.reverse();
        for (int i = 0; i < n; i++){
            int j = i;
            //以空格为界，二次翻转
            while(j < n && res.charAt(j) != ' ')
                j++;
            String temp = res.substring(i,j);
            StringBuffer buffer = new StringBuffer(temp);
            temp = buffer.reverse().toString();
            res.replace(i,j,temp);
            i = j;
        }
        return res.toString();
    }
}

*/
