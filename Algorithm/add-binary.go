package algorithm

/**
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func addBinary(a string, b string) string {
	sa := []rune(a)
	sb := []rune(b)

	al := len(a)
	bl := len(b)
	l := al
	if bl > l {
		l = bl
	}
	l++
	r := make([]rune, l)

	d := 0
	for i := 0; i < l; i++ {
		ac, bc := '0', '0'
		if al-1-i >= 0 {
			ac = sa[al-1-i]
		}
		if bl-1-i >= 0 {
			bc = sb[bl-1-i]
		}
		if ac == bc {
			if d == 1 {
				r[l-1-i] = '1'
			} else {
				r[l-1-i] = '0'
			}
			d = 0
			if ac == '1' {
				d = 1
			}
		} else {
			if d == 1 {
				r[l-1-i] = '0'
			} else {
				r[l-1-i] = '1'
			}
		}

	}

	rr := ""

	bb := false
	for _, c := range r {
		if c == '1' {
			bb = true
		}
		if bb {
			rr = rr + string(c)
		}
	}

	if rr == "" {
		rr = "0"
	}

	return rr
}

/**
自己写的太复杂了。。。

画解算法：67. 二进制求和
灵魂画师牧码
发布于 5 个月前
13.1k 阅读
解题方案
思路
整体思路是将两个字符串较短的用 00 补齐，使得两个字符串长度一致，然后从末尾进行遍历计算，得到最终结果。

本题解中大致思路与上述一致，但由于字符串操作原因，不确定最后的结果是否会多出一位进位，所以会有 2 种处理方式：

第一种，在进行计算时直接拼接字符串，会得到一个反向字符，需要最后再进行翻转
第二种，按照位置给结果字符赋值，最后如果有进位，则在前方进行字符串拼接添加进位
时间复杂度：O(n)O(n)

代码
class Solution {
    public String addBinary(String a, String b) {
        StringBuilder ans = new StringBuilder();
        int ca = 0;
        for(int i = a.length() - 1, j = b.length() - 1;i >= 0 || j >= 0; i--, j--) {
            int sum = ca;
            sum += i >= 0 ? a.charAt(i) - '0' : 0;
            sum += j >= 0 ? b.charAt(j) - '0' : 0;
            ans.append(sum % 2);
            ca = sum / 2;
        }
        ans.append(ca == 1 ? ca : "");
        return ans.reverse().toString();
    }
}

其实思路基本一样
*/
