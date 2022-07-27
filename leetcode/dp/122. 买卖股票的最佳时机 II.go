package dp

func maxProfit(prices []int) int {

	maxProfit := 0

	wallet := -1

	total := len(prices)

	for i := 0; i < total; i++ {

		if wallet == -1 {
			if i == total-1 {
				break
			}
			if prices[i] < prices[i+1] {
				wallet = prices[i]
			}
		} else {
			if prices[i] > wallet {
				maxProfit += (prices[i] - wallet)
				wallet = -1
				i--
			}
		}

	}

	return maxProfit
}

/**
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
执行结果：
通过
显示详情
添加备注

执行用时：
4 ms
, 在所有 Go 提交中击败了
91.33%
的用户
内存消耗：
2.9 MB
, 在所有 Go 提交中击败了
70.29%
的用户
通过测试用例：
200 / 200

自己做的，不过错了一次，忽略了价格为0的情况，离谱


根本没必要动态规划，反而复杂了。。。。
*/

/***
终于碰到个简单题，简单题搞啥动态规划啊，99%

因为交易次数不受限，如果可以把所有的上坡全部收集到，一定是利益最大化的

public int maxProfit(int[] arr) {
        if (arr == null || arr.length <= 1) return 0;

        int ans = 0;
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > arr[i-1]) {  // 卖出有利可图
                ans += (arr[i] - arr[i-1]);
            }
        }

        return ans;
    }
*/
