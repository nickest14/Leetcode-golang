// https://leetcode.com/problems/flip-equivalent-binary-trees/
// Level: Medium

package leetcode

import (
	"leetcode/go/structures"
)

type TreeNode = structures.TreeNode

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == nil && root2 == nil
	}
	if root1.Val != root2.Val {
		return false
	}

	return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) ||
		flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
}
