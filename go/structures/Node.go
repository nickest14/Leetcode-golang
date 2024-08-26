package structures

// ListNode is to define linked list
type Node struct {
	Val      int
	Children []*Node
}

// IntsToNode convert []int to List
func IntsToNode(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	nodes := []*Node{{Val: nums[0], Children: []*Node{}}}
	i := 0
	j := 2
	length := len(nums)

	for j < length {
		if nums[j] != NULL {
			// Create a new node
			tmp := &Node{Val: nums[j], Children: []*Node{}}
			// Attach the new node to the children of the current node
			nodes[i].Children = append(nodes[i].Children, tmp)
			// Add the new node to the nodes list
			nodes = append(nodes, tmp)
		} else {
			// Move to the next parent node
			i++
		}
		j++
	}

	return nodes[0]
}
