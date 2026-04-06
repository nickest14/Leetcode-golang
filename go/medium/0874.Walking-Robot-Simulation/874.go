// https://leetcode.com/problems/walking-robot-simulation/

// Level: Medium

package leetcode

func robotSim(commands []int, obstacles [][]int) int {
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	obstacleSet := make(map[[2]int]struct{}, len(obstacles))
	for _, o := range obstacles {
		obstacleSet[[2]int{o[0], o[1]}] = struct{}{}
	}

	x, y, dir := 0, 0, 0
	ans := 0
	for _, cmd := range commands {
		switch cmd {
		case -1:
			dir = (dir + 1) % 4
		case -2:
			dir = (dir + 3) % 4
		default:
			for range cmd {
				nx := x + dirs[dir][0]
				ny := y + dirs[dir][1]
				if _, hit := obstacleSet[[2]int{nx, ny}]; hit {
					break
				}
				x, y = nx, ny
				if d2 := x*x + y*y; d2 > ans {
					ans = d2
				}
			}
		}
	}
	return ans
}
