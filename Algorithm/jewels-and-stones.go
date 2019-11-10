package algorithm

/**
 给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

示例 1:

输入: J = "aA", S = "aAAbbbb"
输出: 3
示例 2:

输入: J = "z", S = "ZZ"
输出: 0
注意:

S 和 J 最多含有50个字母。
 J 中的字符不重复。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jewels-and-stones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func numJewelsInStones(J string, S string) int {

	var arr [58]byte

	for _, c := range J {
		arr[c-65] = 1
	}

	count := 0

	for _, c := range S {
		if arr[c-65] == 1 {
			count++
		}
	}

	return count
}

/**
执行用时 : 2 ms, 在Jewels and Stones的Java提交中击败了100.00% 的用户。

考虑更高的空间性能，使用byte数组。ASCII码中字母的跨度为65~122，所以定义数组长度为58最节省。

class Solution {
    public int numJewelsInStones(String J, String S) {
        if (S == null || S.isEmpty()) return 0;
        if (J == null || J.isEmpty()) return 0;

        byte[] arr = new byte[58];
        int count = 0;
        for (char ch : J.toCharArray()) {
            arr[ch - 65] = 1;
        }
        for (char ch : S.toCharArray()) {
            if(arr[ch -65] == 1) {
                count++;
            };
        }
        return count;
    }
}
*/

/**
复习关键点：
1.字母的ASCII码是数字且固定
2.利用数组的下标，当成hashmap使用
*/
