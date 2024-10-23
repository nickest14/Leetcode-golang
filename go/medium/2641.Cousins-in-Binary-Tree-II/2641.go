// https://leetcode.com/problems/cousins-in-binary-tree-ii/
// Level: Medium

package leetcode

import (
	"container/list"
	"leetcode/go/structures"
)

type TreeNode = structures.TreeNode

func replaceValueInTree(root *TreeNode) *TreeNode {
	pq := list.New()
	pq.PushBack([2]interface{}{root.Val, root})

	for pq.Len() > 0 {
		n := pq.Len()
		levelSum := 0

		for e := pq.Front(); e != nil; e = e.Next() {
			node := e.Value.([2]interface{})[1].(*TreeNode)
			levelSum += node.Val
		}

		for i := 0; i < n; i++ {
			front := pq.Front()
			localSum := front.Value.([2]interface{})[0].(int)
			node := front.Value.([2]interface{})[1].(*TreeNode)
			pq.Remove(front)

			childSum := 0
			if node.Left != nil {
				childSum += node.Left.Val
			}
			if node.Right != nil {
				childSum += node.Right.Val
			}

			if node.Left != nil {
				pq.PushBack([2]interface{}{childSum, node.Left})
			}
			if node.Right != nil {
				pq.PushBack([2]interface{}{childSum, node.Right})
			}
			node.Val = levelSum - localSum
		}
	}
	return root
}
