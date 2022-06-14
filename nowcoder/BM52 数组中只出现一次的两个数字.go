package nowcoder

func FindNumsAppearOnce(array []int) []int {
	// write code here

	return []int{}
}

/**

这么水的实现？？

方法一：哈希表(推荐使用)
知识点：哈希表

哈希表是一种根据关键码（key）直接访问值（value）的一种数据结构。而这种直接访问意味着只要知道key就能在O(1)O(1)O(1)时间内得到value，因此哈希表常用来统计频率、快速检验某个元素是否出现过等。



方法二：异或运算（扩展思路）
思路：

异或运算满足交换率，且相同的数字作异或会被抵消掉，比如：a⊕b⊕c⊕b⊕c=aa \oplus b \oplus c \oplus b \oplus c=aa⊕b⊕c⊕b⊕c=a，且任何数字与0异或还是原数字，放到这个题目里面所有数字异或运算就会得到a⊕ba \oplus ba⊕b，也即得到了两个只出现一次的数字的异或和。

import java.util.*;
public class Solution {
    public int[] FindNumsAppearOnce (int[] array) {
        int res1 = 0;
        int res2 = 0;
        int temp = 0;
        //遍历数组得到a^b
        for(int i = 0; i < array.length; i++)
            temp ^= array[i];
        int k = 1;
        //找到两个数不相同的第一位
        while((k & temp) == 0)
            k <<= 1;
        for(int i = 0; i < array.length; i++){
            //遍历数组，对每个数分类
            if((k & array[i]) == 0)
                res1 ^= array[i];
            else
                res2 ^= array[i];
        }
        //整理次序
        if(res1 < res2)
            return new int[] {res1, res2};
        else
            return new int[] {res2, res1};
    }
}

*/
