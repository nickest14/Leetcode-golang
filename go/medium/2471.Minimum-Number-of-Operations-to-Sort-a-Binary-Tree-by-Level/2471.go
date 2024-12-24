// https://leetcode.com/problems/cousins-in-binary-tree-ii/
// Level: Medium

package leetcode

import (
	"container/list"
	"leetcode/go/structures"
	"sort"
)

type TreeNode = structures.TreeNode

func minimumOperations(root *TreeNode) int {
	ans := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		curValues := []struct {
			Val int
			Idx int
		}{}
		nextLevel := list.New()

		for queue.Len() > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			curValues = append(curValues, struct {
				Val int
				Idx int
			}{Val: node.Val, Idx: len(curValues)})

			if node.Left != nil {
				nextLevel.PushBack(node.Left)
			}
			if node.Right != nil {
				nextLevel.PushBack(node.Right)
			}
		}

		sort.Slice(curValues, func(i, j int) bool {
			return curValues[i].Val < curValues[j].Val
		})
		idxs := make([]int, len(curValues))
		for i, cv := range curValues {
			idxs[i] = cv.Idx
		}

		for i := 0; i < len(idxs); i++ {
			for i != idxs[i] {
				j := idxs[i]
				idxs[i], idxs[j] = idxs[j], idxs[i]
				ans++
			}
		}

		queue = nextLevel
	}

	return ans
}
