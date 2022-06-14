package nowcoder

func threeSum(num []int) [][]int {
	// write code here

	return [][]int{}
}

/**

描述
给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0？找出数组S中所有满足条件的三元组。

数据范围：0 \le n \le 30000≤n≤3000，数组中各个元素值满足 |val | \le 100∣val∣≤100
空间复杂度：O(n^2)O(n
2
 )，时间复杂度 O(n^2)O(n
2
 )

注意：
三元组（a、b、c）中的元素可以按任意顺序排列。
解集中不能包含重复的三元组。
示例1
输入：
[-10,0,10,20,-10,-40]
复制
返回值：
[[-10,-10,20],[-10,0,10]]
复制
示例2
输入：
[-2,0,1,1,2]
复制
返回值：
[[-2,0,2],[-2,1,1]]
复制
示例3
输入：
[0,0]
复制
返回值：
[]
*/

/**
题目主要信息:
给定一个长度为n的数组，要找出其中所有满足相加等于0的三元组，即数组中所有三个相加为0的数集
三元组内部必须非降序排列，且三元组不能有重复


*/

/**
import java.util.*;
public class Solution {
    public ArrayList<ArrayList<Integer>> threeSum(int[] num) {
        ArrayList<ArrayList<Integer> > res = new ArrayList<ArrayList<Integer>>();
        int n = num.length;
        //不够三元组
        if(n < 3)
            return res;
        //排序
        Arrays.sort(num);
        for(int i = 0; i < n - 2; i++){
            if(i != 0 && num[i] == num[i - 1])
                continue;
            //后续的收尾双指针
            int left = i + 1;
            int right = n - 1;
            //设置当前数的负值为目标
            int target = -num[i];
            while(left < right){
                //双指针指向的二值相加为目标，则可以与num[i]组成0
                if(num[left] + num[right] == target){
                    ArrayList<Integer> temp = new ArrayList<Integer>();
                    temp.add(num[i]);
                    temp.add(num[left]);
                    temp.add(num[right]);
                    res.add(temp);
                    while(left + 1 < right && num[left] == num[left + 1])
                        //去重
                        left++;
                    while(right - 1 > left && num[right] == num[right - 1])
                        //去重
                        right--;
                    //双指针向中间收缩
                    left++;
                    right--;
                }
                //双指针指向的二值相加大于目标，右指针向左
                else if(num[left] + num[right] > target)
                    right--;
                //双指针指向的二值相加小于目标，左指针向右
                else
                    left++;
            }
        }
        return res;
    }
}

*/
