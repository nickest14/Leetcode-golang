// https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/
// Level: Hard

package leetcode

import (
	structures "leetcode/go/structures"
	"strconv"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

type stack struct {
	node  *TreeNode
	depth int
}

func recoverFromPreorder(traversal string) *TreeNode {
	if len(traversal) == 0 {
		return nil
	}

	n := len(traversal)
	i := 0
	var infos [][2]int // (val, depth)

	for i < n {
		depth := 0
		for i < n && traversal[i] == '-' {
			depth++
			i++
		}

		numStart := i
		for i < n && traversal[i] >= '0' && traversal[i] <= '9' {
			i++
		}
		val, _ := strconv.Atoi(traversal[numStart:i])
		infos = append(infos, [2]int{val, depth})
	}

	root := &TreeNode{Val: infos[0][0]}
	stk := []stack{{node: root, depth: 0}}

	for i := 1; i < len(infos); i++ {
		val, depth := infos[i][0], infos[i][1]

		for len(stk) > 0 && stk[len(stk)-1].depth >= depth {
			stk = stk[:len(stk)-1]
		}

		parent := stk[len(stk)-1].node
		node := &TreeNode{Val: val}

		if parent.Left == nil {
			parent.Left = node
		} else {
			parent.Right = node
		}

		stk = append(stk, stack{node: node, depth: depth})
	}

	return root
}
