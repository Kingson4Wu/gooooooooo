package nowcoder

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here

	return false
}

/**

方法一：自顶向下

class Solution {
public:
    map<TreeNode*, int> hs;
    int depth(TreeNode *root) {
        if (!root) return 0;
        if (hs.find(root) != hs.end()) return hs[root];
        int ldep = depth(root->left);
        int rdep = depth(root->right);
        return hs[root] = max(ldep, rdep) + 1;
    }
    bool judge(TreeNode *root) {
        if (!root) return true;
        return abs(hs[root->left] - hs[root->right]) <= 1 &&
        judge(root->left) && judge(root->right);
    }
    bool IsBalanced_Solution(TreeNode* root) {
        depth(root);
        return judge(root);
    }
};

方法二：自底向上
class Solution {
public:
    int depth(TreeNode *root) {
        if (!root) return 0;
        int ldep = depth(root->left);
        if (ldep == -1) return -1;
        int rdep = depth(root->right);
        if (rdep == -1) return -1;
        int sub = abs(ldep - rdep);
        if (sub > 1) return -1;
        return max(ldep, rdep) + 1;
    }
    bool IsBalanced_Solution(TreeNode* root) {
        return depth(root) != -1;
    }
};
*/
