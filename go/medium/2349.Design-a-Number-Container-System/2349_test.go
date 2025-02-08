package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem2349(t *testing.T) {
	obj := Constructor()
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Fine()", obj.Find(10), -1)
	obj.Change(2, 50)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find()", obj.Find(50), 2)
	obj.Change(2, 10)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find()", obj.Find(10), 2)
	obj.Change(1, 10)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find()", obj.Find(10), 1)

}

// ans: list[int] = []
// obj = NumberContainers()
// ans.append(obj.find(10))
// ans.append(obj.change(2, 50))
// ans.append(obj.find(50))
// ans.append(obj.change(2, 10))
// ans.append(obj.find(10))
// ans.append(obj.change(1, 10))
// ans.append(obj.change(3, 10))
// ans.append(obj.find(10))
// print(ans)
