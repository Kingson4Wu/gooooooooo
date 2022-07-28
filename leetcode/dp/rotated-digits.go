package dp

/**
我们称一个数 X 为好数, 如果它的每位数字逐个地被旋转 180 度后，我们仍可以得到一个有效的，且和 X 不同的数。要求每位数字都要被旋转。

如果一个数的每位数字被旋转以后仍然还是一个数字， 则这个数是有效的。0, 1, 和 8 被旋转后仍然是它们自己；2 和 5 可以互相旋转成对方；6 和 9 同理，除了这些以外其他的数字旋转以后都不再是有效的数字。

现在我们有一个正整数 N, 计算从 1 到 N 中有多少个数 X 是好数？

示例:
输入: 10
输出: 4
解释:
在[1, 10]中有四个好数： 2, 5, 6, 9。
注意 1 和 10 不是好数, 因为他们在旋转之后不变。
注意:

N 的取值范围是 [1, 10000]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotated-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//直接排列组合的想法，没想到好的实现方式
/*func rotatedDigits(N int) int {
	num := 0
	count := 0
	for v := N; v > 0; v /= 10 {
		num++
	}
	g := make([]int, num)
	e := make([]int, num)

	temp := N
	for i := 0; i < num; i++ {
		var s int
		if i < num-1 {
			s = temp / ((num - 1 - i) * i)
		} else {
			s = temp
		}

		if s >= 9 {
			g[i] = 7
			e[i] = 7
		} else if s >= 8 {

		}
	}

	return count

}*/

// 动态规划
func rotatedDigits(N int) int {
	count := 0
	dp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if i == 3 || i == 4 || i == 7 ||
			dp[i%10] == 1 || dp[i/10] == 1 {
			dp[i] = 1
		} else if i == 2 || i == 5 || i == 6 || i == 9 ||
			dp[i%10] == 2 || dp[i/10] == 2 {
			// 都不含坏数，且至少包含一个好数，总的就是好数
			dp[i] = 2
			count++
		}
	}
	return count
}

/**
动态规划，时间O(n)，空间O(n)

9以后的每个数n可以拆成a = n % 10（最后一位）和a = n / 10（前面r - 1位）

若a和b中均不含有3、4、7且至少有一个含有2、5、6、9，那么n就是好数

dp数组存储3种值，0：不包含3、4、7的坏数，1：含有3、4、7的坏数，2：好数

通过dp数组可以知晓a和b是否含有3、4、7或2、5、6、9，直接判断出n是否是好数

class Solution {
public:
    int rotatedDigits(int N) {
        int count = 0;
        vector<int> dp(N + 1, 0);
        for (int i = 1; i <= N; i++) {
            if (i == 3 || i == 4 || i == 7 ||
                dp[i % 10] == 1 || dp[i / 10] == 1) {
                dp[i] = 1;
            } else if (i == 2 || i == 5 || i == 6 || i == 9 ||
                dp[i % 10] == 2 || dp[i / 10] == 2) {
                dp[i] = 2;
                count++;
            }
        }
        return count;
    }
};

作者：gremist
链接：https://leetcode-cn.com/problems/rotated-digits/solution/c-san-chong-chang-jian-jie-fa-by-gremist/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

1. 动态规划：保存已经计算过的结果，供后续使用

2. 什么是动态规划（Dynamic Programming）？动态规划的意义是什么？
https://www.zhihu.com/question/23995189

*/
