package nowcoder

func solve(IP string) string {
	// write code here
	return ""
}

/**

描述
编写一个函数来验证输入的字符串是否是有效的 IPv4 或 IPv6 地址

IPv4 地址由十进制数和点来表示，每个地址包含4个十进制数，其范围为 0 - 255， 用(".")分割。比如，172.16.254.1；
同时，IPv4 地址内的数不会以 0 开头。比如，地址 172.16.254.01 是不合法的。

IPv6 地址由8组16进制的数字来表示，每组表示 16 比特。这些组数字通过 (":")分割。比如,  2001:0db8:85a3:0000:0000:8a2e:0370:7334 是一个有效的地址。而且，我们可以加入一些以 0 开头的数字，字母可以使用大写，也可以是小写。所以， 2001:db8:85a3:0:0:8A2E:0370:7334 也是一个有效的 IPv6 address地址 (即，忽略 0 开头，忽略大小写)。

然而，我们不能因为某个组的值为 0，而使用一个空的组，以至于出现 (::) 的情况。 比如， 2001:0db8:85a3::8A2E:0370:7334 是无效的 IPv6 地址。
同时，在 IPv6 地址中，多余的 0 也是不被允许的。比如， 02001:0db8:85a3:0000:0000:8a2e:0370:7334 是无效的。

说明: 你可以认为给定的字符串里没有空格或者其他特殊字符。

数据范围：字符串长度满足 5 \le n \le 505≤n≤50
进阶：空间复杂度 O(n)，时间复杂度 O(n)


示例1
输入：
"172.16.254.1"
复制
返回值：
"IPv4"
复制
说明：
这是一个有效的 IPv4 地址, 所以返回 "IPv4"
示例2
输入：
"2001:0db8:85a3:0:0:8A2E:0370:7334"
复制
返回值：
"IPv6"
复制
说明：
这是一个有效的 IPv6 地址, 所以返回 "IPv6"
示例3
输入：
"256.256.256.256"
复制
返回值：
"Neither"
复制
说明：
这个地址既不是 IPv4 也不是 IPv6 地址
备注：
ip地址的类型，可能为
IPv4,   IPv6,   Neither



方法一：分割字符串比较法（推荐使用）
思路：

我们可以先对IP字符串进行分割，然后依次判断每个分割是否符合要求。

具体做法：

step 1：写一个split函数（或者内置函数）。
step 2：遍历IP字符串，遇到.或者:将其分开储存在一个数组中。
step 3：遍历数组，对于IPv4，需要依次验证分组为4，分割不能缺省，没有前缀0或者其他字符，数字在0-255范围内。
step 4：对于IPv6，需要依次验证分组为8，分割不能缺省，每组不能超过4位，不能出现除数字小大写a-f以外的字符。


import java.util.*;
public class Solution {
    boolean isIPv4 (String IP) {
        if(IP.indexOf('.') == -1){
            return false;
        }
        String[] s = IP.split("\\.");
        //IPv4必定为4组
        if(s.length != 4)
            return false;
        for(int i = 0; i < s.length; i++){
            //不可缺省，有一个分割为零，说明两个点相连
            if(s[i].length() == 0)
                return false;
            //比较数字位数及不为零时不能有前缀零
            if(s[i].length() < 0 || s[i].length() > 3 || (s[i].charAt(0)=='0' && s[i].length() != 1))
                return false;
            int num = 0;
            //遍历每个分割字符串，必须为数字
            for(int j = 0; j < s[i].length(); j++){
                char c = s[i].charAt(j);
                if (c < '0' || c > '9')
                    return false;
            //转化为数字比较，0-255之间
            num = num * 10 + (int)(c - '0');
            if(num < 0 || num > 255)
                return false;
            }
        }
        return true;
    }
    boolean isIPv6 (String IP) {
        if (IP.indexOf(':') == -1) {
            return false;
        }
        String[] s = IP.split(":",-1);
        //IPv6必定为8组
        if(s.length != 8){
            return false;
        }
        for(int i = 0; i < s.length; i++){
            //每个分割不能缺省，不能超过4位
            if(s[i].length() == 0 || s[i].length() > 4){
                return false;
            }
            for(int j = 0; j < s[i].length(); j++){
                //不能出现a-fA-F以外的大小写字符
                char c = s[i].charAt(j);
                boolean expr = (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') ;
                if(!expr){
                    return false;
                }
            }
        }
        return true;
    }
    String solve(String IP) {
        if(isIPv4(IP))
            return "IPv4";
        else if(isIPv6(IP))
            return "IPv6";
        return "Neither";
    }
}


*/
