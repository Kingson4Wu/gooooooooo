package nowcoder

func Serialize(root *TreeNode) string {
	// write code here
	return ""
}

func Deserialize(s string) *TreeNode {
	// write code here
	return nil
}

/**
描述
请实现两个函数，分别用来序列化和反序列化二叉树，不对序列化之后的字符串进行约束，但要求能够根据序列化之后的字符串重新构造出一棵与原二叉树相同的树。

二叉树的序列化(Serialize)是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树等遍历方式来进行修改，序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#）

二叉树的反序列化(Deserialize)是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。


层序序列化(即用函数Serialize转化)如上的二叉树转为"{1,2,3,#,#,6,7}"，再能够调用反序列化(Deserialize)将"{1,2,3,#,#,6,7}"构造成如上的二叉树。

当然你也可以根据满二叉树结点位置的标号规律来序列化，还可以根据先序遍历和中序遍历的结果来序列化。不对序列化之后的字符串进行约束，所以欢迎各种奇思妙想。

*/

/**
方法：前序遍历（推荐使用）
知识点：二叉树递归

思路：

序列化即将二叉树的节点值取出，放入一个字符串中，我们可以按照前序遍历的思路，遍历二叉树每个节点，并将节点值存储在字符串中，我们用‘#’表示空节点，用‘!'表示节点与节点之间的分割。

反序列化即根据给定的字符串，将二叉树重建，因为字符串中的顺序是前序遍历，因此我们重建的时候也是前序遍历，即可还原。


具体做法：

step 1：优先处理序列化，首先空树直接返回“#”，然后调用SerializeFunction函数前序递归遍历二叉树。
SerializeFunction(root, res);
step 2：SerializeFunction函数负责前序递归，根据“根左右”的访问次序，优先访问根节点，遇到空节点在字符串中添加‘#’，遇到非空节点，添加相应节点数字和‘!’，然后依次递归进入左子树，右子树。
//根节点
str.append(root.val).append('!');
//左子树
SerializeFunction(root.left, str);
//右子树
SerializeFunction(root.right, str);
step 3：创建全局变量index表示序列中的下标（C++中直接指针完成）。
step 4：再处理反序列化，读入字符串，如果字符串直接为"#"，就是空树，否则还是调用DeserializeFunction函数前序递归建树。
TreeNode res = DeserializeFunction(str);
step 5：DeserializeFunction函数负责前序递归构建树，遇到‘#’则是空节点，遇到数字则根据感叹号分割，将字符串转换为数字后加入新创建的节点中，依据“根左右”，创建完根节点，然后依次递归进入左子树、右子树创建新节点。
TreeNode root = new TreeNode(val);
......
//反序列化与序列化一致，都是前序
root.left = DeserializeFunction(str);
root.right = DeserializeFunction(str);


import java.util.*;
public class Solution {
    //序列的下标
    public int index = 0;
    //处理序列化的功能函数（递归）
    private void SerializeFunction(TreeNode root, StringBuilder str){
        //如果节点为空，表示左子节点或右子节点为空，用#表示
        if(root == null){
            str.append('#');
            return;
        }
        //根节点
        str.append(root.val).append('!');
        //左子树
        SerializeFunction(root.left, str);
        //右子树
        SerializeFunction(root.right, str);
    }

    public String Serialize(TreeNode root) {
        //处理空树
        if(root == null)
            return "#";
        StringBuilder res = new StringBuilder();
        SerializeFunction(root, res);
        //把str转换成char
        return res.toString();
    }
    //处理反序列化的功能函数（递归）
    private TreeNode DeserializeFunction(String str){
        //到达叶节点时，构建完毕，返回继续构建父节点
        //空节点
        if(str.charAt(index) == '#'){
            index++;
            return null;
        }
        //数字转换
        int val = 0;
        //遇到分隔符或者结尾
        while(str.charAt(index) != '!' && index != str.length()){
            val = val * 10 + ((str.charAt(index)) - '0');
            index++;
        }
        TreeNode root = new TreeNode(val);
        //序列到底了，构建完成
        if(index == str.length())
            return root;
        else
            index++;
        //反序列化与序列化一致，都是前序
        root.left = DeserializeFunction(str);
        root.right = DeserializeFunction(str);
        return root;
    }

    public TreeNode Deserialize(String str) {
        //空序列对应空树
        if(str == "#")
            return null;
        TreeNode res = DeserializeFunction(str);
        return res;
    }
}

*/
