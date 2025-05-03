// https://leetcode.com/problems/push-dominoes/
// Level: Medium

package leetcode

type Pair struct {
	index int
	state byte
}

func pushDominoes(dominoes string) string {
	dom := []byte(dominoes)
	var queue []Pair
	for i, d := range dom {
		if d != '.' {
			queue = append(queue, Pair{index: i, state: d})
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		i, d := p.index, p.state
		if d == 'L' && i > 0 && dom[i-1] == '.' {
			queue = append(queue, Pair{index: i - 1, state: 'L'})
			dom[i-1] = 'L'
		} else if d == 'R' {
			if i+1 < len(dom) && dom[i+1] == '.' {
				if i+2 < len(dom) && dom[i+2] == 'L' {
					if len(queue) > 0 {
						queue = queue[1:]
					}
				} else {
					queue = append(queue, Pair{index: i + 1, state: 'R'})
					dom[i+1] = 'R'
				}
			}
		}
	}
	return string(dom)
}
