package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

//二叉搜索树最近公共祖先
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	if (p.Val < root.Val && q.Val > root.Val) || (p.Val > root.Val && q.Val < root.Val) {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	}

	return nil
}
