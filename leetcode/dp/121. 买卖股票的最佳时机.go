package dp

func MaxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	begin := 0
	end := 0

	min := begin

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[end] {
			end = i
			//begin = min
			if prices[min] < prices[begin] {
				begin = min
			}
			maxProfit = prices[end] - prices[begin]
		} else if prices[i]-prices[min] > maxProfit {
			end = i
			begin = min
			maxProfit = prices[end] - prices[begin]
		} else if prices[i] < prices[min] {
			min = i
		}
	}

	return maxProfit
}

/**
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
多设几个变量，不用动态规划也能完成


执行结果：
通过
显示详情
添加备注

执行用时：
116 ms
, 在所有 Go 提交中击败了
24.51%
的用户
内存消耗：
7.8 MB
, 在所有 Go 提交中击败了
67.80%
的用户
通过测试用例：
211 / 211

输入：
[4,7,1,2,11]
输出：
9
预期结果：
10
*/

/**
答案思路跟我一样，但写得更简单

我们来假设自己来购买股票。随着时间的推移，每天我们都可以选择出售股票与否。那么，假设在第 i 天，如果我们要在今天卖股票，那么我们能赚多少钱呢？

显然，如果我们真的在买卖股票，我们肯定会想：如果我是在历史最低点买的股票就好了！太好了，在题目中，我们只要用一个变量记录一个历史最低价格 minprice，我们就可以假设自己的股票是在那天买的。那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice。

因此，我们只需要遍历价格数组一遍，记录历史最低点，然后在每一天考虑这么一个问题：如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？当考虑完所有天数之时，我们就得到了最好的答案。

public class Solution {
    public int maxProfit(int prices[]) {
        int minprice = Integer.MAX_VALUE;
        int maxprofit = 0;
        for (int i = 0; i < prices.length; i++) {
            if (prices[i] < minprice) {
                minprice = prices[i];
            } else if (prices[i] - minprice > maxprofit) {
                maxprofit = prices[i] - minprice;
            }
        }
        return maxprofit;
    }
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/solution/121-mai-mai-gu-piao-de-zui-jia-shi-ji-by-leetcode-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


*/
