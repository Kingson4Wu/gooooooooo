package nowcoder

func maxLength(arr []int) int {
	// write code here

	return 0
}

/***

知识点1：滑动窗口

滑动窗口是指在数组、字符串、链表等线性结构上的一段，类似一个窗口，而这个窗口可以依次在上述线性结构上从头到尾滑动，且窗口的首尾可以收缩。我们在处理滑动窗口的时候，常用双指针来解决，左指针维护窗口左界，右指针维护窗口右界，二者同方向不同速率移动维持窗口。

知识点2：哈希表

哈希表是一种根据关键码（key）直接访问值（value）的一种数据结构。而这种直接访问意味着只要知道key就能在O(1)O(1)O(1)时间内得到value，因此哈希表常用来统计频率、快速检验某个元素是否出现过等。


import java.util.*;
public class Solution {
    public int maxLength (int[] arr) {
        //哈希表记录窗口内非重复的数字
        HashMap<Integer, Integer> mp = new HashMap<>();
        int res = 0;
        //设置窗口左右边界
        for(int left = 0, right = 0; right < arr.length; right++){
            //窗口右移进入哈希表统计出现次数
            if(mp.containsKey(arr[right]))
                mp.put(arr[right],mp.get(arr[right])+1);
            else
                mp.put(arr[right],1);
            //出现次数大于1，则窗口内有重复
            while(mp.get(arr[right]) > 1)
                //窗口左移，同时减去该数字的出现次数
                mp.put(arr[left],mp.get(arr[left++])-1);
            //维护子数组长度最大值
            res = Math.max(res, right - left + 1);
        }
        return res;
    }
}

*/

/**

给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

数据范围：0\le arr.length \le 10^50≤arr.length≤10
5
 ，0 < arr[i] \le 10^50<arr[i]≤10
5

示例1
输入：
[2,3,4,5]
复制
返回值：
4
复制
说明：
[2,3,4,5]是最长子数组
示例2
输入：
[2,2,3,4,3]
复制
返回值：
3
复制
说明：
[2,3,4]是最长子数组
示例3
输入：
[9]
复制
返回值：
1
复制
示例4
输入：
[1,2,3,1,2,3,2,2]
复制
返回值：
3
复制
说明：
最长子数组为[1,2,3]
示例5
输入：
[2,2,3,4,8,99,3]
复制
返回值：
5
复制
说明：
最长子数组为[2,3,4,8,99]

*/
