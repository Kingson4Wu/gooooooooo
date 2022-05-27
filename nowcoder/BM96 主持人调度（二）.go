package nowcoder

/**
看了答案还是很难理解！！！！
*/

/**

题目的主要信息：
n个活动，有各自的区间
一个主持人不能在相交的区间工作
将相交的区间分成一组，最后组数即是主持人的数量
数字为int型的范围，可能会出现负数
举一反三：
学习完本题的思路你可以解决如下题目：

BM89. 合并区间

BM95. 分糖果问题

方法一：排序+遍历比较（推荐使用）
知识点：贪心思想

贪心思想属于动态规划思想中的一种，其基本原理是找出整体当中给的每个局部子结构的最优解，并且最终将所有的这些局部最优解结合起来形成整体上的一个最优解。

思路：

我们利用贪心思想，什么时候需要的主持人最少？那肯定是所有的区间没有重叠，每个区间首和上一个的区间尾都没有相交的情况，我们就可以让同一位主持人不辞辛劳，一直主持了。但是题目肯定不是这种理想的情况，那我们需要对交叉部分，判断需要增加多少位主持人。

具体做法：

step 1: 利用辅助数组获取单独各个活动开始的时间和结束时间，然后分别开始时间和结束时间进行排序，方便后面判断是否相交。
step 2: 遍历n个活动，如果某个活动开始的时间大于之前活动结束的时候，当前主持人就够了，活动结束时间往后一个。
step 3: 若是出现之前活动结束时间晚于当前活动开始时间的，则需要增加主持人。

复杂度分析：

时间复杂度：O(nlog2n)
​，遍历都是O(n)，sort排序为O(nlog2n)=
空间复杂度：O(n)，辅助空间记录开始时间和结束时间的数组


import java.util.*;
public class Solution {
    public int minmumNumberOfHost (int n, int[][] startEnd) {
        int[] start = new int[n];
        int[] end = new int[n];
        //分别得到活动起始时间
        for(int i = 0; i < n; i++){
            start[i] = startEnd[i][0];
            end[i] = startEnd[i][1];
        }
        //单独排序
        Arrays.sort(start, 0, start.length);
        Arrays.sort(end, 0, end.length);
        int res = 0;
        int j = 0;
        for(int i = 0; i < n; i++){
            //新开始的节目大于上一轮结束的时间，主持人不变
            if(start[i] >= end[j])
                j++;
            else
                //主持人增加
                res++;
        }
        return res;
    }
}




方法二：重载排序+大顶堆（扩展思路）
知识点：优先队列

优先队列即PriorityQueue，是一种内置的机遇堆排序的容器，分为大顶堆与小顶堆，大顶堆的堆顶为最大元素，其余更小的元素在堆下方，小顶堆与其刚好相反。且因为容器内部的次序基于堆排序，因此每次插入元素时间复杂度都是O(log2n)O(log_2n)O(log
2
​
 n)，而每次取出堆顶元素都是直接取出。

具体做法：

step 1：重载sort函数，将开始时间早的活动放在前面，相同情况下再考虑结束时间较早的。
step 2：使用小顶堆辅助，其中堆顶是还未结束的将要最快结束的活动的结束时间。
step 3：首先将int的最小数加入堆中，遍历每一个开始时间，若是堆顶的结束时间小于当前开始时间，可以将其弹出，说明少需要一个主持人。
step 4：除此之外，每次都需要将当前的结束时间需要加入堆中，代表需要一个主持人，最后遍历完成，堆中还有多少元素，就需要多少主持人。

复杂度分析：

时间复杂度：O(nlog2n)
​
 n)，sort排序是O(nlog2n)，一次遍历，循环中维护堆每次O(log2n)
空间复杂度：O(n，堆大小最大为n
*/
