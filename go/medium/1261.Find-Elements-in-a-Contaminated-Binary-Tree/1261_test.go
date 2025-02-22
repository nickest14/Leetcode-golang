package leetcode

import (
	"fmt"
	"leetcode/go/structures"
	"testing"
)

func Test_Problem1261(t *testing.T) {
	root := structures.IntsToTreeNode([]int{-1, structures.NULL, -1})
	obj := Constructor(root)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(1)", obj.Find(1), false)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(2)", obj.Find(2), true)
	fmt.Println()

	root2 := structures.IntsToTreeNode([]int{-1, -1, -1, -1, -1})
	obj2 := Constructor(root2)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(1)", obj2.Find(1), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(3)", obj2.Find(3), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(5)", obj2.Find(5), false)
	fmt.Println()

	root3 := structures.IntsToTreeNode([]int{-1, structures.NULL, -1, -1, structures.NULL, -1})
	obj3 := Constructor(root3)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(2)", obj3.Find(2), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(3)", obj3.Find(3), false)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(4)", obj3.Find(4), false)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find(5)", obj3.Find(5), true)
	fmt.Println()
}
