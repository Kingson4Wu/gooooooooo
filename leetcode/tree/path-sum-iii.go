package tree

/**
给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func pathSum3(root *TreeNode, sum int) int {

	return 0
}

/**
public class Solution437_1 {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
	}

	 public int pathSum(TreeNode root, int sum) {

        if (root == null) {
            return 0;
        }

        return paths(root, sum)
                + pathSum(root.left, sum)
                + pathSum(root.right, sum);
    }

    private int paths(TreeNode root, int sum) {

        if (root == null) {
            return 0;
        }

        int res = 0;
        if (root.val == sum) {
            res += 1;
        }

        res += paths(root.left, sum - root.val);
        res += paths(root.right, sum - root.val);

        return res;
    }

}

作者：li-xin-lei
链接：https://leetcode-cn.com/problems/path-sum-iii/solution/leetcode-437-path-sum-iii-by-li-xin-lei/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

这种时间复杂度高，感觉要用类似动态规划的方法解决


采取了类似于数组的前n项和的思路，比如sum[4] == sum[1]，那么1到4之间的和肯定为0

对于树的话，采取DFS加回溯，每次访问到一个节点，把该节点加入到当前的pathSum中
然后判断是否存在一个之前的前n项和，其值等于pathSum与sum之差
如果有，就说明现在的前n项和，减去之前的前n项和，等于sum，那么也就是说，这两个点之间的路径和，就是sum

最后要注意的是，记得回溯，把路径和弹出去

作者：a380922457
链接：https://leetcode-cn.com/problems/path-sum-iii/solution/liang-chong-fang-fa-jian-dan-yi-dong-ban-ben-by-a3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

class Solution {
    public int pathSum(TreeNode root, int sum) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
        //设置路径和符合条件即res+1的前提（0=pathSum-sum）
        map.put(0, 1);
        return helper(root, map, sum, 0);
    }
    int helper(TreeNode root, HashMap<Integer, Integer> map, int sum, int pathSum){
        int res = 0;
        if(root == null) return 0;
        //将当前所在节点的值加到走过的路径值的和中
        pathSum += root.val;
        //getOrDefault(Object key,V defaultValue)
        // 以上方法为返回指定键（Object key）所映射的值，若无则直接返回所设置的默认值（V defaultValue）
         //累加上到当前节点为止有多少条路径和符合条件（此处若是pathSum-sum==0,则返回1，在map中若存在当前pathSum-sum对应值
        //的key则对应value的值则必不为0，为1或大于1，若无此key则返回方法默认值0）
        res += map.getOrDefault(pathSum - sum, 0);
        //此处是计数到当前节点为止有多少条自上而下的节点路径和等于pathSum，并将其存入map
        // （亦或是更新pathSum对应的路径数，若先前有和值为pathSum的路径则取出其条数先前加上当前的一条）
        map.put(pathSum, map.getOrDefault(pathSum, 0) + 1);
        //往左子树以及右子树依次统计
        // 再加上res-->到当前节点为止可能出现的和值符合pathSum的路径数（统计范围即为头节点到当前节点）
        res = helper(root.left, map, sum, pathSum) + helper(root.right, map, sum, pathSum) + res;
        // 在返回前，将到当前节点为止的和值pathSum的条数计-1，防止影响后面其他未走完路径的统计
        //由于路径和值只能自上而下，所以在当前节点返回前（节点返回条件为下一节点为空，
        // 即为最后节点或者最后节点返回后遍历完依次往上递归返回，返回意味着pathSum到当前节点已自上而下的累加遍历完）
        map.put(pathSum, map.get(pathSum) - 1);
        return res;
    }
}
*/
