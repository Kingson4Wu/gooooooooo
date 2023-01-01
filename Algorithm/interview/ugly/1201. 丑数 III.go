package ugly

/*
*
给你四个整数：n 、a 、b 、c ，请你设计一个算法来找出第 n 个丑数。

丑数是可以被 a 或 b 或 c 整除的 正整数 。

示例 1：

输入：n = 3, a = 2, b = 3, c = 5
输出：4
解释：丑数序列为 2, 3, 4, 5, 6, 8, 9, 10... 其中第 3 个是 4。
示例 2：

输入：n = 4, a = 2, b = 3, c = 4
输出：6
解释：丑数序列为 2, 3, 4, 6, 8, 9, 10, 12... 其中第 4 个是 6。
示例 3：

输入：n = 5, a = 2, b = 11, c = 13
输出：10
解释：丑数序列为 2, 4, 6, 8, 10, 11, 12, 13... 其中第 5 个是 10。
示例 4：

输入：n = 1000000000, a = 2, b = 217983653, c = 336916467
输出：1999999984

提示：

1 <= n, a, b, c <= 10^9
1 <= a * b * c <= 10^18
本题结果在 [1, 2 * 10^9] 的范围内
*/
func NthUglyNumber3(n int, a int, b int, c int) int {
	return nthUglyNumber3(n, a, b, c)
}

func nthUglyNumber3(n int, a int, b int, c int) int {

	var dp int
	pa, pb, pc := 0, 0, 0

	for i := 1; i <= n; i++ {
		ppa, ppb, ppc := pa+a, pb+b, pc+c
		dp = min(min(ppa, ppb), ppc)
		if dp == ppa {
			pa = dp
		}
		if dp == ppb {
			pb = dp
		}
		if dp == ppc {
			pc = dp
		}
	}
	return dp

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
参考丑数II可以做出来,但是不需要dp数组,否则会内存不足
但是时间超出限制,哎

https://mp.weixin.qq.com/s/XXsWwDml_zHiTEFPZtbe3g


这道题的正确解法难度比较大，难点在于你要把一些数学知识和 二分搜索技巧 结合起来才能高效解决这个问题。


首先，我们可以定义一个单调递增的函数f：

f(num, a, b, c)计算[1..num]中，能够整除a或b或c的数字的个数，显然函数f的返回值是随着num的增加而增加的（单调递增）。

题目让我们求第n个能够整除a或b或c的数字是什么，也就是说我们要找到一个最小的num，使得f(num, a, b, c) == n。

这个num就是第n个能够整除a或b或c的数字。

根据 二分查找的实际运用 给出的思路模板，我们得到一个单调函数f，想求参数num的最小值，就可以运用搜索左侧边界的二分查找算法了：

int nthUglyNumber(int n, int a, int b, int c) {
    // 题目说本题结果在 [1, 2 * 10^9] 范围内，
    // 所以就按照这个范围初始化两端都闭的搜索区间
    int left = 1, right = (int) 2e9;
    // 搜索左侧边界的二分搜索
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (f(mid, a, b, c) < n) {
            // [1..mid] 中符合条件的元素个数不足 n，所以目标在右半边
            left = mid + 1;
        } else {
            // [1..mid] 中符合条件的元素个数大于 n，所以目标在左半边
            right = mid - 1;
        }
    }
    return left;
}

// 函数 f 是一个单调函数
// 计算 [1..num] 之间有多少个能够被 a 或 b 或 c 整除的数字
long f(int num, int a, int b, int c) {
    // 下文实现
}
搜索左侧边界的二分搜索代码模板在 二分查找框架详解 中讲过，没啥可说的，关键说一下函数f怎么实现，这里面涉及容斥原理以及最小公因数、最小公倍数的计算方法。

首先，我把[1..num]中能够整除a的数字归为集合A，能够整除b的数字归为集合B，能够整除c的数字归为集合C，那么len(A) = num / a, len(B) = num / b, len(C) = num / c，这个很好理解。

但是f(num, a, b, c)的值肯定不是num / a + num / b + num / c这么简单，因为你注意有些数字可能可以被a, b, c中的两个数或三个数同时整除，如下图：

图片
按照容斥原理，这个集合中的元素应该是：A + B + C - A ∩ B - A ∩ C - B ∩ C + A ∩ B ∩ C。结合上图应该很好理解。

问题来了，A, B, C三个集合的元素个数我们已经算出来了，但如何计算像A ∩ B这种交集的元素个数呢？

其实也很容易想明白，A ∩ B的元素个数就是num / lcm(a, b)，其中lcm是计算最小公倍数（Least Common Multiple）的函数。

类似的，A ∩ B ∩ C的元素个数就是num / lcm(lcm(a, b), c)的值。

现在的问题是，最小公倍数怎么求？

直接记住定理吧：lcm(a, b) = a * b / gcd(a, b)，其中gcd是计算最大公因数（Greatest Common Divisor）的函数。

现在的问题是，最大公因数怎么求？这应该是经典算法了，我们一般叫辗转相除算法（或者欧几里得算法）。

好了，套娃终于套完了，我们可以把上述思路翻译成代码就可以实现f函数，注意本题数据规模比较大，有时候需要用long类型防止int溢出：

// 计算最大公因数（辗转相除/欧几里得算法）
long gcd(long a, long b) {
    if (a < b) {
        // 保证 a > b
        return gcd(b, a);
    }
    if (b == 0) {
        return a;
    }
    return gcd(b, a % b);
}

// 最小公倍数
long lcm(long a, long b) {
    // 最小公倍数就是乘积除以最大公因数
    return a * b / gcd(a, b);
}

// 计算 [1..num] 之间有多少个能够被 a 或 b 或 c 整除的数字
long f(int num, int a, int b, int c) {
    long setA = num / a, setB = num / b, setC = num / c;
    long setAB = num / lcm(a, b);
    long setAC = num / lcm(a, c);
    long setBC = num / lcm(b, c);
    long setABC = num / lcm(lcm(a, b), c);
    // 集合论定理：A + B + C - A ∩ B - A ∩ C - B ∩ C + A ∩ B ∩ C
    return setA + setB + setC - setAB - setAC - setBC + setABC;
}
实现了f函数，结合之前的二分搜索模板，时间复杂度下降到对数级别，即可高效解决这道题目了。

以上就是所有「丑数」相关的题目，用到的知识点有算术基本定理、容斥原理、辗转相除法、链表双指针合并有序链表、二分搜索模板等等。

如果没做过类似的题目可能很难想出来，但只要做过，也就手到擒来了。所以我说这种数学问题属于会者不难，难者不会的类型。
*/
