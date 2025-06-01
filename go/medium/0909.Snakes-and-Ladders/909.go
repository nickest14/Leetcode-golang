// https://leetcode.com/problems/snakes-and-ladders/

// Level: Medium

package leetcode

func snakesAndLadders(board [][]int) int {
	length := len(board)

	intToPos := func(square int) (int, int) {
		r := length - 1 - ((square - 1) / length)
		c := (square - 1) % length
		if (length%2 == 0 && r%2 == 0) || (length%2 != 0 && r%2 != 0) {
			c = length - 1 - c
		}
		return r, c
	}

	type QueueItem struct {
		square int
		moves  int
	}
	queue := []QueueItem{{1, 0}}

	visited := make(map[int]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 1; i <= 6; i++ {
			nextSquare := current.square + i
			if nextSquare > length*length {
				continue
			}

			r, c := intToPos(nextSquare)
			if board[r][c] != -1 {
				nextSquare = board[r][c]
			}
			if nextSquare == length*length {
				return current.moves + 1
			}

			if !visited[nextSquare] {
				visited[nextSquare] = true
				queue = append(queue, QueueItem{nextSquare, current.moves + 1})
			}
		}
	}

	return -1
}
