// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/
// Level: Medium

package leetcode

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	ans := []string{folder[0]}
	for i := 1; i < len(folder); i++ {
		last := ans[len(ans)-1] + "/"
		if !strings.HasPrefix(folder[i], last) {
			ans = append(ans, folder[i])
		}
	}

	return ans
}
