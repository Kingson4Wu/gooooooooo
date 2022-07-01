package tree

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	return []*TreeNode{}
}

/**
给定一棵二叉树 root，返回所有重复的子树。

对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

如果两棵树具有相同的结构和相同的结点值，则它们是重复的。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-duplicate-subtrees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
方法一：深度优先搜索【通过】
思路

序列化二叉树。

class Solution {
    Map<String, Integer> count;
    List<TreeNode> ans;
    public List<TreeNode> findDuplicateSubtrees(TreeNode root) {
        count = new HashMap();
        ans = new ArrayList();
        collect(root);
        return ans;
    }

    public String collect(TreeNode node) {
        if (node == null) return "#";
        String serial = node.val + "," + collect(node.left) + "," + collect(node.right);
        count.put(serial, count.getOrDefault(serial, 0) + 1);
        if (count.get(serial) == 2)
            ans.add(node);
        return serial;
    }
}

作者：LeetCode
链接：https://leetcode.cn/problems/find-duplicate-subtrees/solution/xun-zhao-zhong-fu-de-zi-shu-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：唯一标识符【通过】
思路

假设每棵子树都有一个唯一标识符：只有当两个子树的 id 相同时，认为这两个子树是相同的。

class Solution {
    int t;
    Map<String, Integer> trees;
    Map<Integer, Integer> count;
    List<TreeNode> ans;

    public List<TreeNode> findDuplicateSubtrees(TreeNode root) {
        t = 1;
        trees = new HashMap();
        count = new HashMap();
        ans = new ArrayList();
        lookup(root);
        return ans;
    }

    public int lookup(TreeNode node) {
        if (node == null) return 0;
        String serial = node.val + "," + lookup(node.left) + "," + lookup(node.right);
        int uid = trees.computeIfAbsent(serial, x-> t++);
        count.put(uid, count.getOrDefault(uid, 0) + 1);
        if (count.get(uid) == 2)
            ans.add(node);
        return uid;
    }
}

作者：LeetCode
链接：https://leetcode.cn/problems/find-duplicate-subtrees/solution/xun-zhao-zhong-fu-de-zi-shu-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
